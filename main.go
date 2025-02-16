package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/melnikdev/fiber-grpc/proto"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) GetEvent(ctx context.Context, in *proto.EventRequest) (*proto.EventResponse, error) {
	return &proto.EventResponse{Id: 1, Name: "B-Day", Active: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
