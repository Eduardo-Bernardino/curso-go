package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", servidor.CriateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", servidor.FindUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", servidor.FindUser).Methods(http.MethodGet)

	fmt.Println("Rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
