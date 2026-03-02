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

// GenPass handles HTTP request to generate a password (uses gin).
func (h *passwordHandler) GenPass(c *gin.Context) {
	// TODO: call h.passService.GeneratePassword() and return JSON
	pass, err := h.passService.GeneratePassword()
	if err != nil {
		c.String(http.StatusInternalServerError, "error")
	}
	c.String(http.StatusOK, pass)
}
