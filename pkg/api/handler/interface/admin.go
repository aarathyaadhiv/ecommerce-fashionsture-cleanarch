package interfaces

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	AdminSignUpHandler(c *gin.Context)
	AdminLoginHandler(c *gin.Context)
	BlockUser(c *gin.Context)
	UnblockUser(c *gin.Context)
}
