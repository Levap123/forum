package rest

import (
	"encoding/json"
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
	routes.HandleFunc("/posts/", http.HandlerFunc(h.GetAllPosts))
	routes.HandleFunc("/users", h.User)
	routes.HandleFunc("/users/", h.UserWithId)
	return routes
}

func JSONError(w http.ResponseWriter, err any, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func SendJSON(w http.ResponseWriter, input any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(input)
}
