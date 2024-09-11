package main

import (
	"database/sql"
	"ecom_backend/cmd/api"
	"ecom_backend/config"
	"ecom_backend/db"
	"log"

	_ "github.com/godror/godror"
)

func main() {

	dsn := config.GetDSN()
	db, err := db.NewOracleStorage(dsn)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {

	err := db.Ping() // actually starts the database connection 
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
}
