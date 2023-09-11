package routes

import (
	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.RouterGroup, adminHandler handler.AdminHandler, productHandler handler.ProductHandler,orderHandler handler.OrderHandler,couponHandler handler.CouponHandler) {

	router.POST("/adminLogin", adminHandler.AdminLoginHandler)
	router.Use(middleware.AdminAuthorizationMiddleware)
	{
		router.POST("/adminSignUp", adminHandler.AdminSignUpHandler)
		user := router.Group("/user")
		{
			user.POST("/blockUser/:id", adminHandler.BlockUser)
			user.POST("/unblockUser/:id", adminHandler.UnblockUser)
		}
		product := router.Group("/product")
		{
			product.POST("/add", productHandler.AddProduct)
			product.PATCH("/update", productHandler.UpdateProduct)
			product.DELETE("/delete/:id", productHandler.DeleteProduct)
		}
		category := router.Group("/category")
		{
			category.POST("/add", productHandler.AddCategory)
			category.DELETE("/delete/:id", productHandler.DeleteCategory)
		}
		brand := router.Group("/brand")
		{
			brand.POST("/add", productHandler.AddBrand)
			brand.DELETE("/delete/:id", productHandler.DeleteBrand)
		}
		orders:=router.Group("/orders")
		{
			orders.PATCH("/approval/:id",orderHandler.AdminApproval)
			orders.GET("",orderHandler.ShowOrdersToAdmin)
			orders.GET("/:id",orderHandler.SearchOrder)
		}

		coupon:=router.Group("/coupon")
		{
			coupon.POST("",couponHandler.AddCoupon)
			coupon.PATCH("/expire/:id",couponHandler.ExpireCoupon)
			coupon.PATCH("/block/:id",couponHandler.BlockCoupon)
			coupon.PATCH("/unblock/:id",couponHandler.UnBlockCoupon)
		}
	}
}
