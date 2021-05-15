package main

import (
	"fmt"
	"log"
	"net/http"
	"trackly-backend/app/datastore"
	"trackly-backend/app/security"

	"github.com/gorilla/mux"
)

func handleRequests(port string, db *datastore.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", security.IsAuthorized(handler)).Methods("GET")
	myRouter.Handle("/users", security.IsAuthorized(db.GetAllUsers)).Methods("GET")
	myRouter.Handle("/users/{id}", security.IsAuthorized(db.GetUser)).Methods("GET")
	myRouter.Handle("/users", security.IsAuthorized(db.CreateUser)).Methods("POST")

	fmt.Printf("Serving on port: %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), myRouter))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
}
