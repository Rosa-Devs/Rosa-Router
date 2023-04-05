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
			bytesWritten, err := Conn.WriteTo([]byte(`{"cmd": "1", "id": "pubkey", "tunnel": "None"}`), remote) // id is public key | Defualt tunel data is None
			if err != nil {
				panic(err)
			}

			fmt.Println(bytesWritten, " bytes written")
			time.Sleep(5 * time.Second)
		}

	}()

	listen(Conn, local.String())
}

func listen(conn *net.UDPConn, local string) {
	for {
		fmt.Println("listening")
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[ERROR]", err)
			continue
		}

		fmt.Println("[INCOMING]", string(buffer[0:bytesRead]))
		if string(buffer[0:bytesRead]) == "Hello!" {
			continue
		}

		incoming := string(buffer[0:bytesRead])

		go cli_worker(conn, incoming, local)
	}
}

func cli_worker(conn *net.UDPConn, incoming string, local string) {

	var person manager.Person
	err := json.Unmarshal([]byte(incoming), &person)
	if err != nil {
		panic(err) // debug
	}
	if person.Tunnel != local {
		addr, _ := net.ResolveUDPAddr("udp", person.Tunnel)
		for {
			conn.WriteTo([]byte("Hello!"), addr)
			fmt.Println("sent Hello! to ", person.Tunnel)
			time.Sleep(5 * time.Second)
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
