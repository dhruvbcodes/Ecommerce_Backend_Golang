package api

import (
	"database/sql"
	"log"
	"net/http"
	"ecom_backend/services/user"
	"github.com/gorilla/mux"
)

type APISERVER struct {
	address string
	db      *sql.DB
}

func NewApiServer(address string, db *sql.DB) *APISERVER {
	return &APISERVER{
		address: address,
		db:      db,
	}
}

func (s *APISERVER) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	
	log.Println("Listening on", s.address)
	return http.ListenAndServe(s.address, router)
}
