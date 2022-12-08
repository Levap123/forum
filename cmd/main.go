package main

import (
	"log"

	repository "forum/internal/repository/sqlite3"
	"forum/internal/server"
	"forum/internal/service"
	"forum/internal/transport/rest"
)

func main() {
	db, err := repository.NewDb()
	if err != nil {
		log.Fatal(err) // TODO: FINISH
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := rest.NewHandler(service)
	server := new(server.Server)
	log.Println("server is listening on: http://localhost:8080")
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
