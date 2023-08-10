package main

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Msg string `json:"msg"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	data := HelloResponse{Msg: "Hello world"}
	json.NewEncoder(w).Encode(data)
}
