package routes

import (
	"net/http"

	h "github.com/dwskme/seucy/backend-service/internal/handlers"
	m "github.com/dwskme/seucy/backend-service/internal/middleware"
	s "github.com/dwskme/seucy/backend-service/internal/services"
)

func AuthRoutes(h *h.Handler, tokenService *s.TokenService) {
	http.HandleFunc("/signup", h.SignUp)
	http.HandleFunc("/signin", m.AuthMiddleware(h.SignIn, tokenService))
	http.HandleFunc("/signout", h.SignOut)
}
