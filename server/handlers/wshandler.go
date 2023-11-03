package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")

		allowedOrigins := []string{
			"http://localhost:5173",
			"https://owino-dev.netlify.app",
			"http://http://127.0.0.1:5500/index.html",
		}

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				return true // Allow the connection for this origin
			}
		}

		// Deny the connection for any other origins
		return false
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Received message:", message)

		// get offer from the client and respond with answer
		// handle ice candidate connections

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
