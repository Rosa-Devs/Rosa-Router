package node

import (
	"context"
	"fmt"

	"github.com/Rosa-Devs/Rosa-Router/src/node/api/go/network"
	"github.com/libp2p/go-libp2p/core/host"
	lnet "github.com/libp2p/go-libp2p/core/network" // tihs is not creating conflic with hello protocol in API folder
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Entrypoint
func Start_grpc(port string, host host.Host, ctx context.Context) { // Global func to initialize grpc server from local func
	go start_grpc_server(port, host, ctx)

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
	//log.Println("Node: Message: ", message)
	return &network.Response{Msg: fmt.Sprintf("Hello, %s!", message.Msg)}, nil

}

// Starter
func start_grpc_server(port string, host host.Host, ctx context.Context) error { // local func to start server on go routine
	host.SetStreamHandler("/rosa", func(stream lnet.Stream) {
		// Create a new gRPC server instance

		srv := grpc.NewServer()

		// Register our server implementation
		network.RegisterHelloServiceServer(srv, &server{})

		// Enable reflection for debugging purposes
		reflection.Register(srv)

		// Serve the incoming gRPC requests on the stream
		if err := srv.Serve(); err != nil {
			panic(err)
		}
	})

	// Wait for the host to be closed

}

// func start_grpc_server(port string, host host.Host) error { // local func to start server on go routine
// 	log.Println("Node: Initializing grpc server")
// 	// start the server
// 	srv := grpc.NewServer()
// 	log.Println("Node: GRPC started...")

// 	network.RegisterHelloServiceServer(srv, &server{})

// 	//create lister for serving
// 	listener, err := net.Listen("tcp", ":"+port)
// 	log.Println("Node: GRPC Server started on: ", listener.Addr())
// 	if err != nil {
// 		panic(err)
// 	}

// 	//say server to listen this port
// 	if err := srv.Serve(listener); err != nil {
// 		panic(err)
// 	}

// 	return nil
// }
