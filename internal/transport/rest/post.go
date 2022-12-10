package rest

import (
	"encoding/json"
	"fmt"
	"forum/pkg/webjson"
	"net/http"
)

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	default:
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreatePost(w, r)

	case http.MethodGet:
		GetAllPosts(w, r)
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
