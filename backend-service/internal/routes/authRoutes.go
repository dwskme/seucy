package routes

import (
	"net/http"

	h "github.com/dwskme/seucy/backend-service/internal/handlers"
)

func AuthRoutes(h *h.AuthHandler) {
	http.HandleFunc("/signup", h.SignIn)
	http.HandleFunc("/signin", h.SignIn)
}
