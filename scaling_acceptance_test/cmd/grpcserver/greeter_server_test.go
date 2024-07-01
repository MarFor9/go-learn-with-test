package main_test

import (
	"fmt"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/grpcserver"
	"go-specs-greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
