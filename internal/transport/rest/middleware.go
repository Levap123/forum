package rest

import (
	"context"
	"net/http"
	"time"
)

func (h *Handler) UserIdentity(next http.Handler) http.Handler {
	{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			uuid := cookie.Value
			ctx := context.Background()
			if err != nil {
				ctx = context.WithValue(r.Context(), "id", nil)
			}
			if cookie.Expires.Before(time.Now()) {
				ctx = context.WithValue(r.Context(), "id", nil)
			}
			id, err := h.Service.Auth.GetIdFromSession(uuid)
			if err != nil {
				ctx = context.WithValue(r.Context(), "id", nil)
			}
			if id != 0 {
				ctx = context.WithValue(r.Context(), "id", id)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
