package database

import (
	"fmt"
)

func Prepare_database() {
	fmt.Println("DB: Staring prepare database")
	init_netdb()
	init_node_main()
	fmt.Println("DB: Database prepare DONE!")
}

func init_netdb() {
	fmt.Print("DB: Preparing NODE table")
	db.Exec(`CREATE TABLE IF NOT EXISTS node (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pubkey TEXT,
		ip TEXT,
		port TEXT,
		rating TEXT,
		HS BOOLEAN,
		payload TEXT
	)`)

	fmt.Println(" DONE")

}

func init_node_main() {
	fmt.Print("DB: Preparing NODE_MAIN table")
	db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pubkey TEXT,
		nickname TEXT,
		rating TEXT,
		HS BOOLEAN
	)`)

	//statemen.Exec()
	fmt.Println(" DONE")

}
