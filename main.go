package main

import (
	"log"
	"time"

	"github.com/Rosa-Devs/Rosa-Router/src/node"
	"github.com/Rosa-Devs/Rosa-Router/src/web"
)

// Logger
func init() {

}

// Entrypoint
func main() {

	// init Db
	//database.Start_db()

	log.Println("Main: initialize node...")
	node.Run_node()
	// This for testing GRPC Server worker properly.
	// This code only for TEST not commit and pr to git
	time.Sleep(2 * time.Second)
	//request.Send_test_request()

	// This is main procces, he providing local web api for client
	web.Run_server()
}
