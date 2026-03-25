package handler

import (
	"net/http"

	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type healthHandler struct {
	healthService service.Health
}

type HealthHandler interface {
	Check(c *gin.Context)
}

// check godoc
// @Summary Check health
// @Description Check health
// @Tags    Health
// @Accept  json
// @Produce json
// @Success 200 {object} service.Response
// @Router  /health-check [get]
func NewHealthHandler(healthService service.Health) HealthHandler {
	return &healthHandler{healthService: healthService}
}

func (h *healthHandler) Check(c *gin.Context) {
	resp := h.healthService.Check()
	c.JSON(http.StatusOK, resp)
}
