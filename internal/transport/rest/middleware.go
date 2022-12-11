package rest

import (
	"context"
	"fmt"
	"net/http"
)

func (h *Handler) UserIdentity(next http.Handler) http.Handler {
	{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			ctx := context.Background()
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			fmt.Println(cookie.Expires)
			if cookie.MaxAge < 0 {
				next.ServeHTTP(w, r)
				return
			}
			id, err := h.Service.Auth.GetIdFromSession(cookie.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx = context.WithValue(r.Context(), "id", id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
