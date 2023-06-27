package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Ping handler
// @Description  get pong
// @Tags         ping
// @Success      200  {string}  pong
// @Router       /ping [get]
func (h *Handler) ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
