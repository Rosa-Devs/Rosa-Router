package grrl

import (
	"context"
	"log"

	"github.com/Rosa-Devs/Rosa-Router/src/node/api/go/network"
	"google.golang.org/grpc"
)

func HelloRequest(address string, request *network.Message) (*network.Response, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := network.NewHelloServiceClient(conn)

	response, err := client.SendMessage(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
