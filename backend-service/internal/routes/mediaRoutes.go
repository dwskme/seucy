package routes

import (
	"net/http"

	h "github.com/dwskme/seucy/backend-service/internal/handlers"
)

func MediaRoutes(h *h.NewHandler) {
	http.HandleFunc("/get-preference", h.GetPreferences)
	http.HandleFunc("/add-preference", h.AddPreference)
	http.HandleFunc("/update-preference", h.UpdatePreference)
}
