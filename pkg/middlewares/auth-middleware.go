package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/xhfmvls/restaurant-api/pkg/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
		tokenString := tokenCookie.Value
		
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

		next.ServeHTTP(w, r)
	})
}