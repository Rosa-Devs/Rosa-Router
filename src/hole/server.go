package hole

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/Mihalic2040/Rosa-Router/src/manager"
)

// cache

// type Person struct {
// 	Cmd    string `json:"cmd"`
// 	Id     string `json:"id"`
// 	Tunnel string `json:"tunnel"`
// }

type Cache struct {
	data map[string][]interface{}
}

func NewCache() *Cache {
	return &Cache{data: make(map[string][]interface{})}
}

func (c *Cache) Add(key string, value interface{}) {
	c.data[key] = append(c.data[key], value)
}

func (c *Cache) Get(key string) []interface{} {
	return c.data[key]
}

func (c *Cache) Contains(key string) bool {
	_, ok := c.data[key]
	return ok
}

// Server -- UDP
func Server() {
	localAddress := ":9595"
	if len(os.Args) > 2 {
		localAddress = os.Args[2]
	}

	fmt.Println("HS: UDP Server started: ", localAddress)

	// init new cache

	fmt.Print("HS: Initializing local cache...")

	cache := NewCache()

	// cache.Add("key1", "value2")
	// cache.Add("key2", "value3")
	// cache.Add("key1", "value2")

	fmt.Println(" DONE!")

	//run server

	addr, _ := net.ResolveUDPAddr("udp", localAddress)
	conn, _ := net.ListenUDP("udp", addr)

	for {
		buffer := make([]byte, 1024)
		bytesRead, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		go worker(buffer, bytesRead, remoteAddr, conn, cache)
	}
}

func worker(buffer []byte, bytesRead int, remoteAddr *net.UDPAddr, conn *net.UDPConn, cache *Cache) {
	incoming := string(buffer[0:bytesRead])
	fmt.Println("HS: [INCOMING]", incoming)

	var person manager.Person
	err := json.Unmarshal([]byte(incoming), &person)
	if err != nil {
		panic(err)
	}

	if person.Cmd == "1" {
		if !cache.Contains(person.Id) {
			cache.Add(person.Id, incoming)
			fmt.Println("HS: Adding new node to online DB")
		} else {
			fmt.Println("HS: Node is online now!")
		}

		values := cache.data[person.Id]

		fmt.Println(values)
	}
	//cache.Add(person.Id, incoming)

	//values := cache.data[person.Id]

	//var result string
	//for i, value := range values {
	//	if i > 0 {
	//		result += ","
	//	}
	//	result += fmt.Sprintf("%v", value)
	//}

}
