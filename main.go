package main

import (
	"os"

	"github.com/Mihalic2040/Rosa-Router/src/database"
	"github.com/Mihalic2040/Rosa-Router/src/hole"
	"github.com/Mihalic2040/Rosa-Router/src/web"
)

func main() {
	cmd := os.Args[1]

	// init Db
	database.Start_db()

	switch cmd {
	case "c":
		go hole.Client()
	case "s":
		go hole.Server()
	}
	web.Run_server()
}
