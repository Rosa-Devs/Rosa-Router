package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func create() {
	// Create db folder
	// Define the path to the database directory
	dbDir := "./db"

	// Check if the database directory exists
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		fmt.Println("DB: db folder bot exist!")

		// Create the database directory if it doesn't exist
		err := os.Mkdir(dbDir, 0755)
		if err != nil {
			panic(err)
		}
		fmt.Println("DB: db folder created")
	}

	// create db
	if _, err := os.Stat("./db/database.db"); os.IsNotExist(err) {
		fmt.Println("DB: database.db not exist creating new")
		_, err = os.Create("./db/database.db")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("DB: new database created")
	}
}

func database() {
	var err error
	db, err = sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Other initialization code

}

func Start_db() {
	create()
	fmt.Println("DB: Starting SQLite database")
	database()
	//fmt.Println(" DONE")
	fmt.Println("DB: Starting SQLite database DONE")
	Prepare_database()
}
