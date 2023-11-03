package routes

import (
	// "github.com/owiino/owino-backend-go/internal/middlewares"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	WSRoute(router)

	// Add other routes as needed
	return router
}
