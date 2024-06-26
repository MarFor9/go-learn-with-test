package httpserver

import (
	"fmt"
	"go-specs-greet/domain/interactions"
	"net/http"
)

func Handler(w http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprintf(w, interactions.Greet(name))
}
