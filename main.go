package main

import (
	"database/sql"
	"log"

	"github.com/PanGan21/simple-bank/api"
	db "github.com/PanGan21/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Connect to db
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	// Create a database store
	store := db.NewStore(conn)

	// Create a new server
	server := api.NewServer(store)

	// Start the server
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
