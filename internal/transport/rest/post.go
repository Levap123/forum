package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"forum/pkg/webjson"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

	default:
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userId := strings.TrimPrefix(r.URL.Path, "/posts/")
		if userId == "" {
			h.GetAllPosts(w, r)
			return
		}
	case http.MethodPost:
		h.CreatePost(w, r)
	default:
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
	}
}

type PostInput struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var input PostInput
	userId := r.Context().Value("id")
	if userId == nil {
		webjson.JSONError(w, fmt.Errorf(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		webjson.JSONError(w, fmt.Errorf(http.StatusText(http.StatusInternalServerError)), http.StatusInternalServerError)
		return
	}
	id, err := h.Service.Post.CreatePost(userId.(int), input.Title, input.Body)
	if err != nil {
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]any{"postId": id, "userId": userId})
}

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetPostByUserID(w http.ResponseWriter, r *http.Request, id int) {
}
