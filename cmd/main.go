package main

import (
	"log"

	repository "forum/internal/repository/sqlite3"
	"forum/internal/server"
	"forum/internal/service"
	"forum/internal/transport/rest"
	"forum/pkg/logger"
)

const port = "8080"

func main() {
	db, err := repository.NewDb()
	logger := logger.NewLogger()
	if err != nil {
		logger.Err.Println("Unable to connect database")
		return
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := rest.NewHandler(service, logger)
	server := new(server.Server)
	logger.Info.Printf("server is listening on: http://localhost:%s\n", port)
	log.Printf("server is listening on: http://localhost:%s\n", port)
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
