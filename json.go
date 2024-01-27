package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Printf("responding with 5XX error: %v", msg)
	
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errorResponse{Error: msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal payload: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}