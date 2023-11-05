package routes

import (
	"github.com/streaming-server/handlers"

	"github.com/gorilla/mux"
)

func WSRoute(router *mux.Router) {
	router.HandleFunc("/ws/webtrc", handlers.WSWebRTCHandler)
}
