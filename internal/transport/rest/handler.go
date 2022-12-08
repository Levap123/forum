package rest

import "net/http"

type Handler struct{}

func newHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() http.Handler {
	routes := http.NewServeMux()
	routes.Handle("/posts", UserIdentity(http.HandlerFunc(h.Post)))
	return routes
}
