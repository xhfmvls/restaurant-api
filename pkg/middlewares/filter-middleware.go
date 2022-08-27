package middlewares

import (
	"context"
	"fmt"
	"net/http"
)

const PriceFilterKey ContextKey = "filter"

func PriceFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var minPrice, maxPrice, condQuery string
		minPrice = r.URL.Query().Get("min")
		maxPrice = r.URL.Query().Get("max")

		if minPrice != "" && maxPrice != "" {
			condQuery = fmt.Sprintf("price >= %s AND price <= %s", minPrice, maxPrice)
		} else if minPrice != "" {
			condQuery = fmt.Sprintf("price >= %s", minPrice)
		} else if maxPrice != "" {
			condQuery = fmt.Sprintf("price <= %s", maxPrice)
		} else {
			condQuery = ""
		}

		ctx := context.WithValue(r.Context(), PriceFilterKey, condQuery)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
