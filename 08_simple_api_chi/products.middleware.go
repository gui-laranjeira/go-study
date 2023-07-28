package main

import (
	"context"
	"net/http"
)

func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "custom", "custom")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
