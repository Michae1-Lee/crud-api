package api

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userIDKey contextKey = "userID"

func (h *UserHandler) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}
		headersParts := strings.Split(header, " ")
		if len(headersParts) != 2 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		userId, err := h.authService.ParseToken(headersParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
