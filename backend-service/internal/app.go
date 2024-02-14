package internal

import (
	"log"
	"net/http"
	"time"

	handlers "github.com/dwskme/seucy/backend-service/internal/handlers"
	routes "github.com/dwskme/seucy/backend-service/internal/routes"
	services "github.com/dwskme/seucy/backend-service/internal/services"
	db "github.com/dwskme/seucy/backend-service/internal/utils/db"
)

func App() {
	connectionStr := "postgres://root:root@localhost:5432/seucydb?sslmode=disable"
	secretKey := []byte("thisistheverystupidsecret")

	dbInstance, err := db.InitDB(connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Initialize your services
	authService := services.NewAuthService(dbInstance)
	userService := services.NewUserService(dbInstance)
	tokenService := services.NewTokenService(secretKey, 1*time.Minute)
	mediaService := services.NewMediaService(dbInstance)

	handler := &handlers.NewHandler{
		UserService:  userService,
		TokenService: tokenService,
		AuthService:  authService,
		MediaService: mediaService,
	}

	// Set up routes with middleware
	routes.AuthRoutes(handler, tokenService)
	routes.UserRoutes(handler, tokenService)
	routes.MediaRoutes(handler)

	// Start your server
	http.ListenAndServe(":8080", nil)
}
