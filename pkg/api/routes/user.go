package routes

import (
	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler handler.UserHandler, productHandler handler.ProductHandler, otpHandler handler.OtpHandler) {
	router.POST("/SignUp", userHandler.SignUpHandler)
	router.POST("/login", userHandler.LoginHandler)
	router.POST("/sendOtp", otpHandler.SendOTP)
	router.POST("/verifyOtp", otpHandler.VerifyOtp)
	products := router.Group("/products")
	{
		products.GET("", productHandler.ShowAll)
		products.GET("/:id", productHandler.ShowProduct)
	}
}
