package utils

import (
	"encoding/json"
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
	JsonResponse(w, code, map[string]string{"error": message})
}
