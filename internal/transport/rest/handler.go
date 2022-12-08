package rest

import (
	"net/http"

	"forum/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	routes := http.NewServeMux()
	routes.HandleFunc("/posts/", http.HandlerFunc(h.GetAllPosts))
	routes.HandleFunc("/users/", h.User)
	return routes
}
