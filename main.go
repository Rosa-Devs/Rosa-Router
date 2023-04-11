package main

import (
	"log"
	"os"

	"github.com/Mihalic2040/Rosa-Router/src/node"
	"github.com/Mihalic2040/Rosa-Router/src/web"
)

// Logger
func init() {

}

// Entrypoint
func main() {

	cmd := os.Args[1]

	// init Db
	//database.Start_db()

	switch cmd {
	case "b":
		log.Println("Init bootstrap node...")
	case "c":
		log.Println("Init node...")
		node.Run_node()
	}
	web.Run_server()
}
