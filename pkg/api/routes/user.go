package routes

import (
	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler handler.UserHandler, productHandler handler.ProductHandler, otpHandler handler.OtpHandler, cartHandler handler.CartHandler, orderHandler handler.OrderHandler,paymentHanler handler.PaymentHandler) {
	router.POST("/signup", userHandler.SignUpHandler)
	router.POST("/login", userHandler.LoginHandler)
	router.POST("/sendOtp", otpHandler.SendOTP)
	router.POST("/verifyOtp", otpHandler.VerifyOtp)
	router.GET("/payment/:order_id",paymentHanler.MakePaymentUsingRazorPay)
	router.GET("/payment-success",paymentHanler.VerifyPayment)
	
	password := router.Group("/forgotpassword")
	{
		password.POST("", userHandler.ForgotPassword)
		password.POST("/verify", userHandler.VerifyResetOtp)
		password.Use(middleware.ResetAuthorizationMiddleware)
		{
			password.POST("/reset", userHandler.ResetPassword)
		}
	}
	products := router.Group("/products")
	{
		products.GET("", productHandler.ShowAll)
		products.GET("/:id", productHandler.ShowProduct)
		products.POST("/search",productHandler.SearchProduct)
	}
	
	router.Use(middleware.UserAuthorizationMiddleware)
	{
		profile := router.Group("/userProfile")
		{
			profile.GET("", userHandler.ShowDetails)
			profile.PATCH("", userHandler.UpdateUserDetails)
		}
		address := router.Group("/address")
		{
			address.GET("", userHandler.ShowAddress)
			address.POST("", userHandler.AddAddress)
			address.PATCH("/:id", userHandler.UpdateAddress)
		}
		router.GET("/checkout", userHandler.Checkout)
		cart := router.Group("/cart")
		{
			cart.POST("/:id", cartHandler.AddToCart)
			cart.DELETE("/remove/:id", cartHandler.RemoveFromCart)
			cart.GET("", cartHandler.ShowProductInCart)
			cart.DELETE("",cartHandler.EmptyCart)
		}
		order := router.Group("/orders")
		{
			order.POST("", orderHandler.PlaceOrder)
			order.GET("", orderHandler.ShowOrderHistory)
			order.PATCH("/cancel/:id", orderHandler.CancelOrder)
			order.PATCH("/return/:id",orderHandler.ReturnOrder)
		}
		router.GET("/wallet",orderHandler.GetWallet)
		
	}
}
