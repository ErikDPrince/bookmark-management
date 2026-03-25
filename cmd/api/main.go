package main

import (
	"log"

	"github.com/ErikDPrince/bookmark-management/internal/api"
)

// @title           Bookmark Management API
// @description     Bookmark Management API
// @version         1.0
// @host           localhost:8080
// @basePath       /api
// @schemes        http
// @securityDefinitions.apikey BearerAuth
// @in             header
// @name           Authorization
func main() {

	cfg, err := api.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Đặt tên khác package `api` để tránh shadow và dễ đọc.
	engine := api.New(cfg)
	engine.Start()
}
