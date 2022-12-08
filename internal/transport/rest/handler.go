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
	routes.Handle("/posts", UserIdentity(http.HandlerFunc(h.Post)))
	routes.HandleFunc("/users", h.User)
	return routes
}
