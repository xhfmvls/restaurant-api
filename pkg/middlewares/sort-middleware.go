package middlewares

import (
	"context"
	"net/http"
	"strings"
)

type ContextKey string

const SortKey ContextKey = "sort"

func Sorting(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sortType := r.URL.Query().Get("sort")
		if sortType == "" {
			sortType = "id-asc"
		}
		switch sortType {
		case
			"id-asc",
			"price-asc",
			"price-desc",
			"name-asc",
			"name-desc":
			sortType = strings.ReplaceAll(sortType, "-", " ")
		default:
			panic("Sort Type not Valid")
		}
		ctx := context.WithValue(r.Context(), SortKey, sortType)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}