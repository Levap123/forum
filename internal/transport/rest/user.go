package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"forum/internal/entities"
	"forum/pkg/webjson"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	if idStr == "" {
		h.GetAllUsers(w, r)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		webjson.JSONError(w, fmt.Errorf("Not found"), http.StatusNotFound)
		return
	}
	h.GetUserById(w, r, id)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request, userId int) {
	var user entities.User
	user, err := h.Service.User.GetUserById(userId)
	if err != nil {
		webjson.JSONError(w, err, http.StatusNotFound)
		return
	}
	webjson.SendJSON(w, user)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []entities.User
	users, err := h.Service.User.GetAllUsers()
	if err != nil {
		webjson.JSONError(w, err, http.StatusInternalServerError)
	}
	webjson.SendJSON(w, users)
}
