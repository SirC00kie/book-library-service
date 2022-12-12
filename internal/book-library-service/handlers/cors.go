package handlers

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodDelete,
			http.MethodPost,
		},
		AllowedHeaders:     []string{},
		ExposedHeaders:     []string{},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
	})
	return c
}
