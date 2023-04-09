package database

import (
	"fmt"
	"log"

	"github.com/Mihalic2040/Rosa-Router/src/serialization"
)

func AddNewNode(Node string) {

	add(Node)
}

func add(Node string) {
	node, _ := serialization.DeserializeNode(Node)

	//fmt.Printf("node: %+v\n", node)

	count := FindNodeById(Node)

	//fmt.Println(count)
	fmt.Println("DB: Addding NODE", "Pubkey:", node.Pubkey, "Ip:", node.Ip, "Port:", node.Port)

	if count > 0 {
		// Update the existing record
		_, err := db.Exec("UPDATE node SET ip = ?, port = ?, rating = ?, hs = ?, hsport = ? WHERE pubkey = ?",
			node.Ip, node.Port, node.Rating, node.Hs, node.HsPort, node.Pubkey)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// Insert a new record
		_, err := db.Exec("INSERT INTO node (pubkey, ip, port, rating, hs, hsport) VALUES (?, ?, ?, ?, ?, ?)",
			node.Pubkey, node.Ip, node.Port, node.Rating, node.Hs, node.HsPort)
		if err != nil {
			log.Fatal(err)
		}
	}
}
