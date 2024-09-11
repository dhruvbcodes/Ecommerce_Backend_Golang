package main

import (
	"ecom_backend/cmd/api"
	"ecom_backend/config"
	"ecom_backend/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
