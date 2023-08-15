package interfaces

import "github.com/gin-gonic/gin"

type ProductHandler interface {
	AddProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	AddCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	ShowAll(c *gin.Context)
	ShowProduct(c *gin.Context)
	AddBrand(c *gin.Context)
	DeleteBrand(c *gin.Context)
}
