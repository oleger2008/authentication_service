package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/oleger2008/authentication_service/pkg/user/v1"
)

const grpcPort = 50051

type Server struct {
	desc.UnimplementedUserV1Server
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create response name = %s", req.Name)
	log.Printf("Create response email = %s", req.Email)
	log.Printf("Create response password = %s", req.Password)
	log.Printf("Create response password_confirm = %s", req.PasswordConfirm)
	log.Printf("Create response role = %d", req.Role)

	return &desc.CreateResponse{
		Id: 42,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get response id = %d", req.Id)

	return &desc.GetResponse{
		Id:    req.Id,
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Role:  1,
	}, nil
}

func (s *Server) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	log.Printf("Update response id = %d", req.Id)
	return &empty.Empty{}, nil
}

func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	log.Printf("Delete response id = %d", req.Id)
	return &empty.Empty{}, nil
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
