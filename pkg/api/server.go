package http

import (
	"github.com/gin-gonic/gin"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/middleware"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handler.UserHandler, adminHandler handler.AdminHandler, productHandler handler.ProductHandler, otpHandler handler.OtpHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, otpHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
