package handler

import (
	"net/http"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type healthHandler struct {
	healthService service.Health
}

func NewHealthHandler(healthService service.Health) HealthHandler {
	return &healthHandler{healthService: healthService}
}

func (h *healthHandler) Check(c *gin.Context) {