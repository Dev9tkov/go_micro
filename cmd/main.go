package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Dev9tkov/go_micro/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())
	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User name: %s", req.GetName())
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("User new name %s", req.GetName())
	return nil, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User for delete %d", req.GetId())
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
