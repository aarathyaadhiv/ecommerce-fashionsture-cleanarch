package interfaces

import "github.com/gin-gonic/gin"


type OrderHandler interface{
	PlaceOrder(c *gin.Context)
	ShowOrderHistory(c *gin.Context)
	CancelOrder(c *gin.Context)
	AdminApproval(c *gin.Context)
	ReturnOrder(c *gin.Context)
	ShowOrdersToAdmin(c *gin.Context)
	SearchOrder(c *gin.Context)
	FilterOrderByApproval(c *gin.Context)
	FilterOrderByPaymentStatus(c *gin.Context)
	GetWallet(c *gin.Context)
}