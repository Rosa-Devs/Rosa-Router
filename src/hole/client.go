package hole

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Mihalic2040/Rosa-Router/src/manager"
)

var Conn *net.UDPConn

// Client --
func Client() {
	register()
}

func register() {
	signalAddress := os.Args[2]

	localAddress := ":9595" // default port
	if len(os.Args) > 3 {
		localAddress = os.Args[3]
	}

	remote, _ := net.ResolveUDPAddr("udp", signalAddress)
	local, _ := net.ResolveUDPAddr("udp", localAddress)
	Conn, _ := net.ListenUDP("udp", local)
	go func() {
		for {
			bytesWritten, err := Conn.WriteTo([]byte(`{
				"cmd" : "1",
				"pubkey": "puubkey123",
				"ip": "127.0.0.1",
				"port": "8080",
				"rating": "5",
				"hs": "false",
				"hsport": "8443"
			}
			`), remote) // id is public key | Defualt tunel data is None
			if err != nil {
				panic(err)
			}
			fmt.Println("CLI: send to HS:", bytesWritten, " bytes")
			time.Sleep(5 * time.Second)
		}

	}()

	listen(Conn, local.String())
}

func listen(conn *net.UDPConn, local string) {
	for {
		fmt.Println("CLI: listening on", local)
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("CLI: [ERROR]", err)
			continue
		}

		fmt.Println("CLI: [INCOMING]", string(buffer[0:bytesRead]))

		incoming := string(buffer[0:bytesRead])

		go cli_worker(conn, incoming, local)
	}
}

func cli_worker(conn *net.UDPConn, incoming string, local string) {
	//cmd handling

	//adding new node to db
	var nodes []manager.Node
	if err := json.Unmarshal([]byte(incoming), &nodes); err != nil {
		panic(err)
	}

	fmt.Printf("Unmarshaled %d nodes:\n", len(nodes))
	for _, node := range nodes {
		Tunnel := "[" + node.Ip + "]" + ":" + node.Port
		fmt.Println(Tunnel)
		if Tunnel != local {
			addr_rmt, _ := net.ResolveUDPAddr("udp", Tunnel)
			for {
				conn.WriteTo([]byte("Hello!"), addr_rmt)
				fmt.Println("sent Hello! to ", Tunnel)
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func Get_node(id string) {
	signalAddress := os.Args[2]

	//localAddress := ":9595" // default port
	//if len(os.Args) > 3 {
	//	localAddress = os.Args[3]
	//}

	remote, _ := net.ResolveUDPAddr("udp", signalAddress)
	//local, _ := net.ResolveUDPAddr("udp", localAddress)
	conn := Conn

	p := manager.New_connection_in{Cmd: "2", Id: id}
	// Convert the person object to JSON
	jsonBytes, err := json.Marshal(p)

	bytesWritten, err := conn.WriteTo(jsonBytes, remote) // id is public key | Defualt tunel data is None
	if err != nil {
		panic(err)
	}

	fmt.Println(bytesWritten, " bytes written")
}
