package internal

import (
	"log"
	"net/http"

	"github.com/dwskme/seucy/backend-service/internal/handlers"
	"github.com/dwskme/seucy/backend-service/internal/routes"
	"github.com/dwskme/seucy/backend-service/internal/services"
	"github.com/dwskme/seucy/backend-service/internal/utils/db"
)

func App() {
	// TODO:Change connectionString to be dynamic
	connectionStr := "postgres://root:root@localhost:5432/seucydb?sslmode=disable"
	dbInstance, err := db.InitDB(connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()
	// Initialize your services
	authService := services.NewAuthService(dbInstance)
	userService := services.NewUserService(dbInstance)
	tokenService := services.NewTokenService([]byte("sertyuiolkjdcbnm"))

	// Initialize your handler
	authHandler := &handlers.AuthHandler{
		UserService:  userService,
		TokenService: tokenService,
		AuthService:  authService,
	}

	// Set up routes
	routes.AuthRoutes(authHandler)

	// Start your server
	http.ListenAndServe(":8080", nil)
}
