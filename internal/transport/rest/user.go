package rest

import (
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
		if err != nil {
		}
		GetUserById(w, r, id)
	}
	
}

func GetUserById(w http.ResponseWriter, r *http.Request, userId int) {
}
