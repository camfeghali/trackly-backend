package main

import (
	"fmt"
	"log"
	"net/http"
	"trackly-backend/app/datastore"

	"github.com/gorilla/mux"
)

func handleRequests(port string, db *datastore.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handler).Methods("GET")
	myRouter.HandleFunc("/users", db.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", db.GetUser).Methods("GET")
	myRouter.HandleFunc("/users", db.CreateUser).Methods("POST")

	fmt.Printf("Serving on port: %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), myRouter))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
}
