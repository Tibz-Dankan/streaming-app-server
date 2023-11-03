package httpserver

import (
	"log"
	"net/http"

	"github.com/streaming-server/routes"

	"github.com/rs/cors"
)

func Start() {
	router := routes.AppRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	http.Handle("/", handler)

	log.Println("Starting http server up on 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
