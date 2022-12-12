package rest

import (
	"net/http"

	"forum/internal/service"
	"forum/pkg/logger"
)

type Handler struct {
	Service *service.Service
	Logger  *logger.Logger
}

func NewHandler(service *service.Service, logger *logger.Logger) *Handler {
	return &Handler{
		Service: service,
		Logger:  logger,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	routes := http.NewServeMux()
	routes.HandleFunc("/auth/sign-in", h.SignIn)
	routes.HandleFunc("/auth/sign-up", h.SignUp)
	routes.Handle("/auth/sign-out", h.UserIdentity(http.HandlerFunc(h.SignOut)))

	routes.HandleFunc("/users/", h.User)

	routes.HandleFunc("/posts/", h.GetPosts)
	routes.Handle("/users/posts/", h.UserIdentity(http.HandlerFunc(h.Posts)))

	routes.Handle("/posts/action", h.UserIdentity())
	return routes
}
