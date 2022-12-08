package rest

import (
	"encoding/json"
	"net/http"
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
	}
}
