package node

import "log"

func Start_grpc() { // Global func to initialize grpc server from local func
	err := start_grpc_server()
	if err != nil {
		log.Println("Node: failed to start grpc server try to run as root") // IDK ssome times it lags and need to run as root
		log.Println("Node: grpc server error: ", err)                       // Showing the error to log
	}
}

func start_grpc_server() error {
	log.Println("Node: Initializing grpc server")

	return nil
}
