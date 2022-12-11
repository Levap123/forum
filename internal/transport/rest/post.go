package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/pkg/webjson"
)

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	default:
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	default:
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")
	if userId == nil {
		// 401 no auth
	}
	err := json.NewEncoder(w).Encode(userId)
	if err != nil {
		// 500?
	}
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
}
