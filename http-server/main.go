package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewPostgresPlayerStore()}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
