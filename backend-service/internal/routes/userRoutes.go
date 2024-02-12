package routes

import (
	"net/http"

	h "github.com/dwskme/seucy/backend-service/internal/handlers"
	m "github.com/dwskme/seucy/backend-service/internal/middleware"
	s "github.com/dwskme/seucy/backend-service/internal/services"
)

func UserRoutes(u *h.Handler, tokenService *s.TokenService) {
	http.HandleFunc("/test", m.AuthMiddleware(u.EnglishHandler, tokenService))
}
