package rest

import (
	"encoding/json"
	"net/http"

	"forum/internal/entities"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
}

type signUpInput struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input signUpInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}
	user := entities.User{
		Email:    input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	id, err := h.Service.Auth.CreateUser(user)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}
	SendJSON(w, map[string]int{"userId": id})
}
