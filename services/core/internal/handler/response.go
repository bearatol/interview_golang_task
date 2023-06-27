package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type successResponse struct {
	Message string `json:"message"`
}
type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func newErrorResponse(c *gin.Context, log *zap.SugaredLogger, statusCode int, err error, message string) {
	if err == nil {
		err = fmt.Errorf("")
	}
	log.Errorf("%s, error: %s", message, err)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Message: message,
		Error:   fmt.Sprintf("%s", err),
	})
}

func newSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, successResponse{
		Message: "success",
	})
}
