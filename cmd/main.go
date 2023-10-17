package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/oleger2008/authentication_service/pkg/user/v1"
)

const grpcPort = 50051

type Server struct {
	desc.UnimplementedUserV1Server
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterUserV1Server(server, &Server{})

	log.Printf("Server is listening at %v", listener.Addr())

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
