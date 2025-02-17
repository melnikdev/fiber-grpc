package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/melnikdev/fiber-grpc/config"
	"github.com/melnikdev/fiber-grpc/database"
	"github.com/melnikdev/fiber-grpc/internal/model"
	proto "github.com/melnikdev/fiber-grpc/proto"
	"google.golang.org/grpc"
)

type ServerGRPC struct {
	App    *grpc.Server
	Config *config.Config
	DB     database.Database
}
type server struct {
	proto.UnimplementedAddServiceServer
	DB database.Database
}

func (s *server) GetEvent(ctx context.Context, in *proto.EventRequest) (*proto.EventResponse, error) {
	var event model.Event

	s.DB.GetDb().First(&event, 1)

	fmt.Println(event)
	return &proto.EventResponse{Id: int64(event.ID), Name: event.Name, Active: event.Active}, nil
}

func NewServerGRPC(c *config.Config, db database.Database) *ServerGRPC {
	return &ServerGRPC{
		App:    grpc.NewServer(),
		Config: c,
		DB:     db,
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
	proto.RegisterAddServiceServer(s.App, &server{DB: s.DB})
}
