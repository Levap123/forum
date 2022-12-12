package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"forum/internal/entities"
	"forum/pkg/errors"
	"forum/pkg/webjson"
)

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userIdStr := strings.TrimPrefix(r.URL.Path, "/posts/")
		if userIdStr == "" {
			h.GetAllPosts(w, r)
			return
		}
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			webjson.JSONError(w, errors.WebFail(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		h.GetPostByPostId(w, r, userId)
	default:
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (h *Handler) Posts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		postIdStr := strings.TrimPrefix(r.URL.Path, "/users/posts/")
		if postIdStr == "" {
			webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
			return
		}
		postId, err := strconv.Atoi(postIdStr)
		if err != nil {
			webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
			return
		}
		h.GetPostsByUserId(w, r, postId)
	case http.MethodPost:
		h.CreatePost(w, r)
	default:
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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
		h.Logger.Err.Println("uuid is not correct")
		webjson.JSONError(w, errors.WebFail(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := h.Service.Post.CreatePost(userId.(int), input.Title, input.Body)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusUnauthorized), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]any{"postId": id, "userId": userId})
}

func (h *Handler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []entities.Post
	posts, err := h.Service.Post.GetAllPosts()
	if posts == nil {
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, posts)
}

func (h *Handler) GetPostsByUserId(w http.ResponseWriter, r *http.Request, id int) {
	var posts []entities.Post
	posts, err := h.Service.Post.GetAllUsersPosts(id)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if posts == nil {
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	webjson.SendJSON(w, posts)
}

func (h *Handler) GetPostByPostId(w http.ResponseWriter, r *http.Request, id int) {
	var post entities.Post
	post, err := h.Service.Post.GetPostByPostId(id)
	if post.Id == 0 {
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, post)
}
