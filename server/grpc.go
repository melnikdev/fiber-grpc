package server

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/melnikdev/fiber-grpc/config"
	proto "github.com/melnikdev/fiber-grpc/proto"
	"google.golang.org/grpc"
)

type ServerGRPC struct {
	App    *grpc.Server
	Config *config.Config
}
type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) GetEvent(ctx context.Context, in *proto.EventRequest) (*proto.EventResponse, error) {
	return &proto.EventResponse{Id: 1, Name: "B-Day", Active: true}, nil
}

func NewServerGRPC(c *config.Config) *ServerGRPC {
	return &ServerGRPC{
		App:    grpc.NewServer(),
		Config: c,
	}
}

func (s *ServerGRPC) Start() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(s.Config.Server.Port))
	if err != nil {
		panic(err)
	}
	s.registerService()
	log.Printf("server listening at %v", lis.Addr())
	if err := s.App.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *ServerGRPC) registerService() {
	proto.RegisterAddServiceServer(s.App, &server{})
}
