package database

import (
	"fmt"
	"log"

	"github.com/Mihalic2040/Rosa-Router/src/serialization"
)

func FindNodeById(Node string) int {
	fmt.Println("DB: Addding NODE: ", Node)
	count := find(Node)
	return count
}

func find(Node string) int {
	node, _ := serialization.DeserializeNode(Node)

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM node WHERE pubkey = ?", node.Pubkey).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
