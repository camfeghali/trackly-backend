package main

import (
	"fmt"
	"log"
	"net/http"
	"trackly-backend/app/datastore"
	"trackly-backend/app/security"

	"github.com/gorilla/mux"
)

var authorizer *security.Security = &security.Security{AuthorizationEnabled: GetConfig("prod").AUTHORIZATION_ENABLED}

func handleRequests(port int, db *datastore.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/", authorizer.IsAuthorized(handler)).Methods("GET")
	myRouter.Handle("/users", authorizer.IsAuthorized(db.GetAllUsers)).Methods("GET")
	myRouter.Handle("/users/{id}", authorizer.IsAuthorized(db.GetUser)).Methods("GET")
	myRouter.HandleFunc("/users", db.CreateUser).Methods("POST")
	myRouter.HandleFunc("/login", db.Login).Methods("POST")

	fmt.Printf("Serving on port: %v\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), myRouter))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root")
}
