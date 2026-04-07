package main

import (
	"log"

	"github.com/ErikDPrince/bookmark-management/internal/api"
	"github.com/ErikDPrince/bookmark-management/pkg/logger"
)

// @title           Bookmark Management API
// @description     Default listen port is 8080 (APP_PORT). Swagger Try it out uses @host below — run the server on the same port (unset APP_PORT or APP_PORT=8080).
// @version         1.0
// @host            localhost:8080
// @basePath        /
// @schemes         http
// @securityDefinitions.apikey BearerAuth
// @in             header
// @name           Authorization
func main() {
	logger.SetLogLevel()
	cfg, err := api.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	engine, err := api.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}
}
