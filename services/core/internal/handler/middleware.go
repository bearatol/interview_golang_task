package handler

import (
	"net/http"
	"strings"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

func timeoutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func (h *Handler) TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(mapping.TimeoutConnect),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func (h *Handler) userIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, h.log, http.StatusUnauthorized, nil, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, h.log, http.StatusUnauthorized, nil, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, h.log, http.StatusUnauthorized, nil, "token is empty")
			return
		}

		err := h.user.UserCheck(h.ctx, headerParts[1])
		if err != nil {
			h.log.Error(err)
			newErrorResponse(c, h.log, http.StatusUnauthorized, err, "invalid token")
			return
		}

		c.AddParam("token", headerParts[1])
		c.Next()
	}

}
