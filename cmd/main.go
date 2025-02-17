package main

import (
	_ "github.com/lib/pq"
	"github.com/melnikdev/fiber-grpc/config"
	"github.com/melnikdev/fiber-grpc/database"
	"github.com/melnikdev/fiber-grpc/internal/model"
	"github.com/melnikdev/fiber-grpc/server"
)

func main() {
	config := config.MustLoad()
	db := database.NewPostgresDatabase(config)

	db.GetDb().AutoMigrate(&model.Event{}, &model.User{})

	app := server.NewServerGRPC(config, db)
	app.Start()
}
