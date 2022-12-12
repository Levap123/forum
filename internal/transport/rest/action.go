package rest

import (
	"encoding/json"
	"net/http"

	"forum/pkg/errors"
	"forum/pkg/webjson"
)

type VoteInput struct {
	Action string `json:"action,omitempty"`
	PostId int    `json:"post_id,omitempty"`
}

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		webjson.JSONError(w, errors.WebFail(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var input VoteInput
	id := r.Context().Value("id")
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		webjson.JSONError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.Service.Action.VotePost(id.(int), input.PostId, input.Action); err != nil {
		h.Logger.Err.Println(err)
		webjson.JSONError(w, errors.WebFail(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	likes, _, err := h.Service.Action.GetPostVotes(input.PostId)
	if err != nil {
		h.Logger.Err.Println(err)
		webjson.JSONError(w, errors.WebFail(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	webjson.SendJSON(w, map[string]any{"votes": likes})
}
