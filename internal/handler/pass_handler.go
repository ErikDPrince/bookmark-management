package handler

import (
	"net/http"

	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type passwordHandler struct {
	passService service.Password
}

type Password interface {
	GenPass(c *gin.Context)
}

func NewPassword(svc service.Password) Password {
	return &passwordHandler{passService: svc}
}

// GenPass handles HTTP request to generate a password.
// @Summary Generate a password
// @Tags Password
// @Produce json
// @Success 200 {object} map[string]string
// @Router  /password [get]
func (h *passwordHandler) GenPass(c *gin.Context) {
	pass, err := h.passService.GeneratePassword()
	if err != nil {
		c.String(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
