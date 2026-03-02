package main

import (
	"github.com/ErikDPrince/bookmark-management/internal/api"
)

func main() {
	api := api.New()
	api.Start()
}
