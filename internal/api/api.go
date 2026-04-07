package api

import (
	"fmt"

	_ "github.com/ErikDPrince/bookmark-management/docs"

	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/repository"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	redispkg "github.com/ErikDPrince/bookmark-management/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Engine interface {
	Start() error
}

type engine struct {
	app         *gin.Engine
	cfg         *Config
	redisClient *redis.Client
}

// New builds the HTTP engine: Redis client, Gin router, and all routes.
func New(cfg *Config) (Engine, error) {
	redisClient, err := redispkg.NewClient("")
	if err != nil {
		return nil, err
	}

	e := &engine{
		app:         gin.New(),
		cfg:         cfg,
		redisClient: redisClient,
	}
	e.initRoutes()
	return e, nil
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%s", e.cfg.AppPort))
}

// initRoutes wires HTTP routes (read top-to-bottom like a route table).
func (e *engine) initRoutes() {
	// Password
	passSvc := service.NewPassWordService()
	passH := handler.NewPassword(passSvc)
	e.app.GET("/password", passH.GenPass)

	// Shorten URL (matches Swagger: POST /v1/links/shorten)
	repo := repository.NewURLStorage(e.redisClient)
	shortenSvc := service.NewShortenURLService(repo, service.NewCodeGenerator())
	shortenH := handler.NewShortenURLHandler(shortenSvc)
	e.app.POST("/v1/links/shorten", shortenH.ShortenURL)
	e.app.GET("/v1/links/:code", shortenH.GetURL)

	// Health
	healthSvc := service.NewHealthService(e.cfg.ServiceName, e.cfg.InstanceID)
	healthH := handler.NewHealthHandler(healthSvc)
	e.app.GET("/health-check", healthH.Check)

	// Swagger UI
	e.app.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.DocExpansion("none"),
	))
}
