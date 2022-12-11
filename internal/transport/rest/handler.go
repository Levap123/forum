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
	routes.Handle("/auth/sign-out", h.UserIdentity(http.HandlerFunc(h.SignOut)))

	routes.HandleFunc("/posts/", h.GetPosts)
	routes.HandleFunc("/users/", h.User)
	routes.Handle("/users/posts/", h.UserIdentity(http.HandlerFunc(h.Posts)))
	return routes
}
