package internal

import (
	"net/http"

	"github.com/dwskme/seucy/backend-service/internal/routes"
	"github.com/dwskme/seucy/backend-service/internal/utils/db"
)

func App() {
	db.InitDB()
	defer db.CloseDB()
	routes.InitializeRoutes()
	http.ListenAndServe(":8080", nil)
}
