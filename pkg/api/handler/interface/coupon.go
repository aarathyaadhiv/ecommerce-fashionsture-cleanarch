package interfaces

import "github.com/gin-gonic/gin"

type CouponHandler interface{
	AddCoupon(c *gin.Context)
	ExpireCoupon(c *gin.Context)
	BlockCoupon(c *gin.Context)
	UnBlockCoupon(c *gin.Context)
	GetCoupon(c *gin.Context)
}