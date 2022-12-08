package rest

import (
	"context"
	"net/http"
	"time"
)

func UserIdentity(next http.Handler) http.Handler {
	{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			ctx := context.Background()
			if err != nil {
				ctx = context.WithValue(r.Context(), "id", nil)
			}
			if cookie.Expires.Before(time.Now()) {
				ctx = context.WithValue(r.Context(), "id", nil)
			}
			id := 0 // TODO: get id from uuid
			ctx = context.WithValue(r.Context(), "id", id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
