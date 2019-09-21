package main

import (
	"fmt"
	"net/http"

	"github.com/ConradPacesa/k8s-backend/api-gateway/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/items/all", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/new", handlers.NewItem).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Print(err)
	}
}
