package interfaces

import "github.com/gin-gonic/gin"



type CartHandler interface{
	AddToCart(c *gin.Context)
	RemoveFromCart(c *gin.Context)
	ShowProductInCart(c *gin.Context)
}