package middleware

import (
	"net/http"
	"time"

	services "github.com/dwskme/seucy/backend-service/internal/services"
)

func AuthMiddleware(handler http.HandlerFunc, tokenService *services.TokenService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		///check for signin case
		if r.URL.Path == "/signin" && r.Method == http.MethodPost {
			identifier := r.FormValue("identifier")
			token, err := tokenService.GenerateToken(identifier)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			// Attach the token to the response in an HTTP-only cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token,
				Expires:  time.Now().Add(tokenService.TokenExpires),
				HttpOnly: true,
				SameSite: http.SameSiteStrictMode,
			})
			handler(w, r)
			return
		}
		// TODO:change this to use token form header later after working on client
		// token, err := tokenService.GetTokenFromBearerHeader(r)
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Validate the token
		_, valid, err := tokenService.ValidateToken(cookie.Value)
		if err != nil || !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}
