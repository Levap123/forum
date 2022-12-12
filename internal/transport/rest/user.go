package rest

import (
	"net/http"
	"strconv"
	"strings"

	"forum/internal/entities"
	"forum/pkg/errors"
	"forum/pkg/webjson"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	if idStr == "" {
		h.GetAllUsers(w, r)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	h.GetUserById(w, r, id)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request, userId int) {
	var user entities.User
	user, err := h.Service.User.GetUserById(userId)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusNotFound)
		return
	}
	webjson.SendJSON(w, user)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []entities.User
	users, err := h.Service.User.GetAllUsers()
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	webjson.SendJSON(w, users)
}
