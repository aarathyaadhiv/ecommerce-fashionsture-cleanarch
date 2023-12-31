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
	ShowCategory(c *gin.Context)
	ShowBrand(c *gin.Context)
	SearchProduct(c *gin.Context)
	UpdateCategory(c *gin.Context)
	UpdateBrand(c *gin.Context)
	GetProductsToAdmin(c *gin.Context)
}
