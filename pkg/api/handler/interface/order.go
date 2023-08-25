package interfaces

import "github.com/gin-gonic/gin"




type OrderHandler interface{
	PlaceOrder(c *gin.Context)
	ShowOrderHistory(c *gin.Context)
}