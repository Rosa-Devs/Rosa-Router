package request

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Rosa-Devs/Rosa-Router/src/node/api/go/network"
	"github.com/Rosa-Devs/Rosa-Router/src/node/grrl"
)

func Send_test_request() {
	for {
		// generate random nuber for request id
		rand.Seed(time.Now().UnixNano())

		// Generate a random integer between 100000 and 999999
		randomNumber := rand.Intn(900000) + 100000

		// Convert the integer to a string
		randomString := strconv.Itoa(randomNumber)

		request := &network.Message{Msg: randomString}

		response, err := grrl.HelloRequest("localhost:50051", request)
		if err != nil {
			log.Fatalf("could not call RPC method: %v", err)
		}

		log.Printf("Response: %v", response)
		time.Sleep(5 * time.Second)
	}

}
