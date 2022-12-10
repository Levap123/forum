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
	routes.HandleFunc("/auth/sign-in", h.SignIn)
	routes.HandleFunc("/auth/sign-up", h.SignUp)
	routes.Handle("/posts/", h.UserIdentity(http.HandlerFunc(h.GetAllPosts)))
	routes.HandleFunc("/users/", h.User)
	return routes
}
