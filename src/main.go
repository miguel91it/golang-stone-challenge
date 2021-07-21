package main

import (
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

var db Storage

func init() {

	// inicia o repositorioinmemory
	db = NewStorage()

	InitAccounts()

}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}/balance", GetAccountBalance).Methods("GET")
	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	router.HandleFunc("/transfers", GetTransfers).Methods("GET")
	router.HandleFunc("/transfers", MakeTransfer).Methods("POST")

	router.HandleFunc("/login", LoginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
