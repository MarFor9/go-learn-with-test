package main

import (
	"go-specs-greet/adapters/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &grpcserver.GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
