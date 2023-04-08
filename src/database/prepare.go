package database

import (
	"fmt"
	"log"
)

func Prepare_database() {
	fmt.Println("DB: Staring prepare database")
	init_netdb()
	init_node_main()
	fmt.Println("DB: Database prepare DONE!")
}

func init_netdb() {
	fmt.Print("DB: Preparing NODE table")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS node (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pubkey TEXT,
		ip TEXT,
		port TEXT,
		rating TEXT,
		hs TEXT,
		hsport TEXT
	)`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" DONE")

}

func init_node_main() {
	fmt.Print("DB: Preparing NODE_MAIN table")
	db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pubkey TEXT,
		rating TEXT,
		HS BOOLEAN
	)`)

	//statemen.Exec()
	fmt.Println(" DONE")

}
