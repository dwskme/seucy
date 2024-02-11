package middleware

import (
	"net/http"

	services "github.com/dwskme/seucy/backend-service/internal/services"
)

func AuthRequired(handler http.HandlerFunc, tokenService *services.TokenService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := tokenService.ExtractTokenFromHeader(r.Header.Get("Authorization"))
		if err != nil || tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// Validate the token using your TokenService
		valid, err := tokenService.ValidateToken(tokenString)
		if err != nil || !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		handler(w, r)
	}
}
