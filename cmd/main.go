package main

import (
	"log"

	repository "forum/internal/repository/sqlite3"
	"forum/internal/service"
	"forum/internal/transport/rest"
	app "forum/internal/transport/server"
)

func main() {
	db, err := repository.NewDb()
	if err != nil {
		log.Fatal(err) // TODO: FINISH
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := rest.NewHandler(service)
	server := new(app.Server)
	log.Println("Listen")
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
