package node

import (
	"log"

	"github.com/libp2p/go-libp2p"

	mplex "github.com/libp2p/go-libp2p/p2p/muxer/mplex"
	quic "github.com/libp2p/go-libp2p/p2p/transport/quic"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
)

func Run_node() {
	node()
}

func node() {
	host, err := libp2p.New(
		//address
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
}
