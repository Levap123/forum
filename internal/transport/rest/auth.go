package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"forum/internal/entities"
	"forum/pkg/webjson"
)

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	var input signInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		webjson.JSONError(w, err, http.StatusBadRequest)
	}
	uuid, err := h.Service.CreateSession(input.Email, input.Password)
	if err != nil {
		webjson.JSONError(w, err, http.StatusUnauthorized)
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   uuid,
		Expires: time.Now().Add(24 * time.Hour),
	})
}

type signUpInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, fmt.Errorf("Method not allowed"), http.StatusMethodNotAllowed)
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
		webjson.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]int{"userId": id})
}
