package http

import (
	"github.com/gin-gonic/gin"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/middleware"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handler.UserHandler, adminHandler handler.AdminHandler, productHandler handler.ProductHandler, otpHandler handler.OtpHandler, cartHandler handler.CartHandler, orderHandler handler.OrderHandler, paymentHandler handler.PaymentHandler, couponHandler handler.CouponHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("templates/*.html")

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, otpHandler, cartHandler, orderHandler, paymentHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, orderHandler, couponHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3008")
}
     