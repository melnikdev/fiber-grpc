package main

import (
	"github.com/melnikdev/fiber-grpc/config"
	"github.com/melnikdev/fiber-grpc/server"
)

func main() {
	config := config.MustLoad()
	app := server.NewServerGRPC(config)
	app.Start()
}
