package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func JsonResponse(w http.ResponseWriter, code int, toJson interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toJson)
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	fmt.Printf("Error: %v\n", message)
	JsonResponse(w, code, map[string]string{"error": message})
}
