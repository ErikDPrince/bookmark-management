package api

import (
	_ "github.com/ErikDPrince/bookmark-management/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Start() error
}

type api struct {
	app *gin.Engine
	cfg *Config
}

func New(cfg *Config) Engine {
	a := &api{
		app: gin.New(),
		cfg: cfg,
	}
	a.registerEP()
	a.registerHealthEP()
	a.registerSwaggerEP()

	return a
}

func (a *api) Start() error {
	return a.app.Run(":" + a.cfg.AppPort)
}

func (a *api) registerEP() {
	passSvc := service.NewPassWordService()
	passHandler := handler.NewPassword(passSvc)
	a.app.GET("/password", passHandler.GenPass)
}

func (a *api) registerHealthEP() {
	healthSvc := service.NewHealthService(a.cfg.ServiceName, a.cfg.InstanceID)
	healthHandler := handler.NewHealthHandler(healthSvc)
	a.app.GET("/health-check", healthHandler.Check)
}

func (a *api) registerSwaggerEP() {
	a.app.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.DocExpansion("none"),
	))
}
