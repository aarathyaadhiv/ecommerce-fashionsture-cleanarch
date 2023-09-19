package routes

import (
	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler handler.AdminHandler, productHandler handler.ProductHandler, orderHandler handler.OrderHandler, couponHandler handler.CouponHandler) {

	router.POST("/login", adminHandler.AdminLoginHandler)
	router.Use(middleware.AdminAuthorizationMiddleware)
	{
		router.POST("/signup", adminHandler.AdminSignUpHandler)
		router.GET("", adminHandler.AdminHome)
		router.GET("/dashboard", adminHandler.Dashboard)
		router.GET("/salesreport", adminHandler.SalesReport)
		user := router.Group("/user")
		{
			user.GET("", adminHandler.ListUsers)
			user.PATCH("/block/:id", adminHandler.BlockUser)
			user.PATCH("/unblock/:id", adminHandler.UnblockUser)
		}
		product := router.Group("/product")
		{
			product.GET("",productHandler.GetProductsToAdmin)
			product.POST("", productHandler.AddProduct)
			product.PATCH("/:id", productHandler.UpdateProduct)
			product.DELETE("/:id", productHandler.DeleteProduct)
		}
		category := router.Group("/category")
		{
			category.GET("",productHandler.ShowCategory)
			category.POST("", productHandler.AddCategory)
			category.DELETE("/:id", productHandler.DeleteCategory)
			category.PATCH("/:id", productHandler.UpdateCategory)
		}
		brand := router.Group("/brand")
		{
			brand.GET("",productHandler.ShowBrand)
			brand.POST("", productHandler.AddBrand)
			brand.DELETE("/:id", productHandler.DeleteBrand)
			brand.PATCH("/:id", productHandler.UpdateBrand)
		}
		orders := router.Group("/orders")
		{
			orders.PATCH("/approval/:id", orderHandler.AdminApproval)
			orders.GET("", orderHandler.ShowOrdersToAdmin)
			orders.GET("/:id", orderHandler.SearchOrder)
			
		}

		coupon := router.Group("/coupon")
		{
			coupon.GET("", couponHandler.GetCoupon)
			coupon.POST("", couponHandler.AddCoupon)
			coupon.PATCH("/expire/:id", couponHandler.ExpireCoupon)
			coupon.PATCH("/block/:id", couponHandler.BlockCoupon)
			coupon.PATCH("/unblock/:id", couponHandler.UnBlockCoupon)
		}
	}
}
