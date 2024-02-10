package routes

import (
	"net/http"

	"github.com/dwskme/seucy/backend-service/internal/handlers"
)

func InitializeRoutes() {
	http.HandleFunc("/create-user", handlers.CreateUserHandler)
	http.HandleFunc("/get-user", handlers.GetUserHandler)
}
