package main

import (
	"database/sql"
	"log"

	"github.com/PanGan21/simple-bank/api"
	db "github.com/PanGan21/simple-bank/db/sqlc"
	"github.com/PanGan21/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	// Loas configuration and env variables
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load conf:", err)
	}

	// Connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	// Create a database store
	store := db.NewStore(conn)

	// Create a new server
	server := api.NewServer(store)

	// Start the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
