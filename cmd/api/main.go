package main

import (
	"log"

	"github.com/ErikDPrince/bookmark-management/internal/api"
)

func main() {

	cfg, err := api.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Đặt tên khác package `api` để tránh shadow và dễ đọc.
	engine := api.New(cfg)
	engine.Start()
}
