package handler

import (
	"context"
	"time"

	"go.uber.org/zap"

	_ "github.com/bearatol/interview_golang_task/sevices/core/api/swagger"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/service"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service interface{}

type Handler struct {
	ctx     context.Context
	log     *zap.SugaredLogger
	user    User
	product Products
	price   Prices
}

// @title Swagger Router API
// @version 1.0
// @description API Router

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type 'Bearer TOKEN' to correctly set the API Key. Example: Bearer <access_token>

// @host localhost:6001
// @BasePath /api/v1

func NewHandler(
	ctx context.Context,
	log *zap.SugaredLogger,
	serv *service.Service,
) *Handler {
	return &Handler{
		ctx:     ctx,
		log:     log,
		user:    serv,
		product: serv,
		price:   serv,
	}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = mapping.MaxMsgSize

	router.Use(ginzap.Ginzap(h.log.Desugar(), time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(h.log.Desugar(), true))
	router.Use(cors.Default())
	router.Use(h.TimeoutMiddleware())
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", h.ping)

		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		users := v1.Group("/users")
		{
			users.POST("/regis", h.UserRegistration)
			users.GET("/auth", h.UserAuth)

			access := users.Group("/")
			access.Use(h.userIdentity())
			{
				access.GET("/", h.UserGet)
				access.PUT("/", h.UserUpdate)
				access.DELETE("/", h.UserDelete)
			}
		}

		products := v1.Group("/products")
		products.Use(h.userIdentity())
		{
			products.GET("/", h.ProductGet)
			products.POST("/", h.ProductCreate)
			products.PUT("/", h.ProductUpdate)
			products.DELETE("/", h.ProductDelete)

			prices := products.Group("/prices")
			{
				prices.GET("/", h.PricesGet)
				prices.GET("/:filename", h.PricesGetFile)
				prices.POST("/", h.PriceCreate)
				prices.DELETE("/", h.PriceDelete)
			}
		}
	}

	return router
}
