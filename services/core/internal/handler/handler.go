package handler

import (
	"context"
	"time"

	"go.uber.org/zap"

	_ "github.com/bearatol/interview_golang_task/sevices/core/api/swagger"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/middleware"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service interface{}

type Handler struct {
	ctx  context.Context
	log  *zap.SugaredLogger
	serv Service
}

// @title Swagger Router API
// @version 1.0
// @description API Router

// @host localhost:6001
// @BasePath /api/v1

func NewHandler(ctx context.Context, log *zap.SugaredLogger, serv Service) *Handler {
	return &Handler{
		ctx:  ctx,
		log:  log,
		serv: serv,
	}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = mapping.MaxMsgSize

	router.Use(ginzap.Ginzap(h.log.Desugar(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(h.log.Desugar(), true))
	router.Use(cors.Default())
	router.Use(middleware.TimeoutMiddleware())

	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// router.Use(sessions.Sessions("hr_session", store))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", h.ping)

		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}
