package node

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/libp2p/go-libp2p"

	dht "github.com/libp2p/go-libp2p-kad-dht" // idk this some bug with importing DHT i need to import it 2x to use func called .New()
	"github.com/multiformats/go-multiaddr"

	kademlia "github.com/libp2p/go-libp2p-kad-dht"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	mplex "github.com/libp2p/go-libp2p/p2p/muxer/mplex"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

func Run_node() {
	get_environment()
	node()
}

var port string

func get_environment() string {
	node_mod := os.Args[1]

	switch node_mod {
	case "b":
		fmt.Println()
		port = "9595"
		return "/ip4/0.0.0.0/tcp/9563/" // default bootstrap nodes
	}
	return "/ip4/0.0.0.0/tcp/0/"
}

func node() {
	ctx := context.Background()

	host, err := libp2p.New(
		//address
		libp2p.ListenAddrStrings(get_environment()), // use for use dynamic port // to use own port type it to /0

		//inint transport
		libp2p.Transport(quic.NewTransport),
		libp2p.Transport(tcp.NewTCPTransport),
		//inint multiplexing
		libp2p.Muxer("/mplex/6.7.0", mplex.DefaultTransport),
		libp2p.EnableHolePunching(),
		libp2p.EnableNATService(),
	)
	if err != nil {
		log.Println("Node: error creating host: ", err)
		log.Println("Node: host id is: ", host.ID())
	}

	log.Println("Node: host id is: ", host.ID())
	log.Println("Node: links: ", host.Addrs())

	log.Println("Node: Starting GRPC server...")

	Start_grpc(port, host, ctx) // Entry point of GRPC server // port here

	//DHT configuration

	KademliaDHT := initDHT(ctx, host)

	bootstrap(ctx, KademliaDHT, host) // dont need to go run routine it go routine automatically
	go bootDHT(ctx, KademliaDHT, host)

}

func initDHT(ctx context.Context, host host.Host) *dht.IpfsDHT {
	log.Println("Node: DHT: initializing hash table...")
	dhtOpts := []dht.Option{
		dht.ProtocolPrefix("/rosa"),
	}

	KademliaDHT, err := kademlia.New(ctx, host, dhtOpts...)
	if err != nil {
		log.Println("Node: DHT: init failed with error: ", err)
	}

	log.Println("Node: DHT: table created!")

	return KademliaDHT

}

func bootstrap(ctx context.Context, kademlia *kademlia.IpfsDHT, host host.Host) {
	log.Println("Node: DHT: bootstrap thread started!")
	if err := kademlia.Bootstrap(ctx); err != nil {
		log.Println("Node: DHT: bootstrap thread failed: ", err)
	}
}

func bootDHT(ctx context.Context, KademliaDHT *dht.IpfsDHT, host host.Host) {

	// this is test
	// in release bootstrap nodes will be dynamically updated and we need to update this list from DB
	// how we solve bootstrap problem read on README file
	bootnodes := []string{
		"/ip4/127.0.0.1/tcp/9563",
	}

	log.Println("Node: DHT: booting from:", bootnodes)

	// we stating cycle for connecting to bootnodes
	// connect to each bootstrap node in a loop
	for _, addr := range bootnodes {
		pi, err := peer.AddrInfoFromP2pAddr(multiaddr.StringCast(addr))
		if err != nil {
			log.Printf("Node: DHT: failed to parse bootstrap node address %s: %s\n", addr, err)
			continue
		}

		log.Printf("Node: DHT: connecting to bootstrap node %s\n", pi.ID)
		if err := host.Connect(ctx, *pi); err != nil {
			log.Printf("Node: DHT: failed to connect to bootstrap node %s: %s\n", pi.ID, err)
			continue
		}

		log.Printf("Node: DHT: connected to bootstrap node %s\n", pi.ID)
	}

	// Bootstrap the DHT
	if err := KademliaDHT.Bootstrap(ctx); err != nil {
		log.Println("Node: DHT: bootstrap thread failed: ", err)
	}

}
