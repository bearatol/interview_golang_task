package middleware

import (
	"net/http"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func timeoutResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(mapping.TimeoutConnect),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}
