package hole

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// cache

type Person struct {
	Cmd string `json:"cmd"`
	Id  string `json:"id"`
}

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

// Server -- UDP
func Server() {
	localAddress := ":9595"
	if len(os.Args) > 2 {
		localAddress = os.Args[2]
	}

	fmt.Println("UDP Server started: ", localAddress)

	// init new cache

	fmt.Print("Initializing local cache...")

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
	fmt.Println("[INCOMING]", incoming)

	var person Person
	err := json.Unmarshal([]byte(incoming), &person)
	if err != nil {
		panic(err)
	}

	cache.Add(person.Id, incoming)

	values := cache.data[person.Id]

	var result string
	for i, value := range values {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%v", value)
	}

}
