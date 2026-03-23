package api

import (
	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Start() error
}

type api struct {
	app *gin.Engine
}

func New() Engine {
	return &api{app: gin.Default()}
}

func (a *api) Start() error {
	return a.app.Run(":8080")
}

func (a *api) registerEP() {
	passSvc := service.NewPassWordService()
	passHandler := handler.NewPassword(passSvc)
	a.app.GET("/password", passHandler.GenPass)
}
