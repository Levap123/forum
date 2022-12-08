package rest

import (
	"log"
	"net/http"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	log.Println(123123)
}
