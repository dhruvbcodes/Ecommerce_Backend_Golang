package db

import (
	"database/sql"
	"log"
	_ "github.com/godror/godror"
)

func NewOracleStorage(dsn string) (*sql.DB, error) {
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}