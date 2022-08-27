package middlewares

import (
	"context"
	"net/http"
)

const SearchKey ContextKey = "search"

func Search(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		searchVal := r.URL.Query().Get("name")
		ctx := context.WithValue(r.Context(), SearchKey, searchVal)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}