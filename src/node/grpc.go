package node

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Rosa-Devs/Rosa-Router/src/node/api/go/network"
	"google.golang.org/grpc"
)

// Entrypoint
func Start_grpc() { // Global func to initialize grpc server from local func
	go start_grpc_server()
	// if err != nil {
	// 	log.Println("Node: failed to start grpc server try to run as root") // IDK ssome times it lags and need to run as root
	// 	log.Println("Node: grpc server error: ", err)                       // Showing the error to log
	// }
}

// Code

// Services
type server struct {
	network.UnimplementedHelloServiceServer
}

// Handlers and logical components
func (s *server) SendMessage(ctx context.Context, message *network.Message) (*network.Response, error) {
	// Implement the SendMessage method here
	return &network.Response{Msg: fmt.Sprintf("Hello, %s!", message.Msg)}, nil
}

// Starter
func start_grpc_server() error { // local func to start server on go routine
	log.Println("Node: Initializing grpc server")
	// start the server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	network.RegisterHelloServiceServer(srv, &server{})
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
	return nil
}
