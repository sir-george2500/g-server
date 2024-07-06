package main

import (
	"log"
	"net/http"
)

func responWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responWithJson(w, code, errResponse{
		Error: msg,
	})
}

// handle respon with payload
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	responWithJson(w, 200, struct{}{})
}
