package middlewares

import (
	"context"
	"net/http"
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/xhfmvls/restaurant-api/pkg/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

const IdKey ContextKey = "id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("Token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(tokenCookie.Value, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace(tokenCookie.Value, "Bearer ", "", -1)

		claims := models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, &claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), IdKey, claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
