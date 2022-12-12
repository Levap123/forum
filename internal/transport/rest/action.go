package rest

import (
	"net/http"

	"forum/pkg/errors"
	"forum/pkg/webjson"
)

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	
}
