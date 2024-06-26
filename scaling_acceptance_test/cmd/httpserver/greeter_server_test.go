package main

import (
	"fmt"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/httpserver"
	"go-specs-greet/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)
	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
