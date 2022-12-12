package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/internal/entities"
	"forum/pkg/errors"
	"forum/pkg/webjson"
)

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		webjson.JSONError(w, err, http.StatusBadRequest)
		return
	}
	uuid, err := h.Service.CreateSession(input.Email, input.Password)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, fmt.Errorf("Invalid credentials"), http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  uuid,
		Path:   "/",
		MaxAge: int(24 * time.Hour),
	})

	webjson.SendJSON(w, map[string]any{"sessionId": uuid})
}

type signUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, errors.WebFail(http.StatusNotFound), http.StatusMethodNotAllowed)
		return
	}
	var input signUpInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		webjson.JSONError(w, err, http.StatusBadRequest)
		return
	}
	user := entities.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	id, err := h.Service.Auth.CreateUser(user)
	if err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]any{"userId": id})
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusMethodNotAllowed)
		return
	}
	userId := r.Context().Value("id")
	if userId == nil {
		webjson.JSONError(w, errors.WebFail(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if err := h.Service.Auth.DeleteSession(userId.(int)); err != nil {
		h.Logger.Err.Println(err.Error())
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	webjson.SendJSON(w, map[string]any{"message": "deleted"})
}
