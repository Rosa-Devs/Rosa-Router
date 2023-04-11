package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func server(localAddress string) {

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register handler functions for different routes
	mux.HandleFunc("/api/v0.1/new_connection", new_connection)

	// Start a new Goroutine to listen for incoming connections
	go func() {
		if err := http.ListenAndServe(localAddress, mux); err != nil {
			fmt.Println(err)
		}
	}()

	// Block the main Goroutine from exiting
	select {}
}

func Run_server() {
	localAddress := ":9595" // default port
	if len(os.Args) > 3 {
		localAddress = os.Args[3]
	}
	log.Println("WEB: Starting web api on ", localAddress)
	server(localAddress)
}
