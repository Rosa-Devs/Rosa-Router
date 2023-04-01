package main

import (
	"os"

	"github.com/Mihalic2040/Rosa-Router/src/hole"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "c":
		hole.Client()
	case "s":
		hole.Server()
	}
}
