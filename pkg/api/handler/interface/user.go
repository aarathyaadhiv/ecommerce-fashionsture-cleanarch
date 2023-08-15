package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	SignUpHandler(c *gin.Context)
	LoginHandler(c *gin.Context)
	// LoginOtpHandler(c *gin.Context)
	// FindAll(c *gin.Context)
	// FindByID(c *gin.Context)
	// Save(c *gin.Context)
	// Delete(c *gin.Context)
}
