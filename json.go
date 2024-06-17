package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	s, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Print(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(s)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Println("500 error",msg)
	}

	type errResponse struct{
		Error string `json:"error"`
	}

	responseWithJson(w,code,errResponse{
		Error: msg,
	})
}


