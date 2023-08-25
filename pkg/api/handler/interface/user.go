package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	SignUpHandler(c *gin.Context)
	LoginHandler(c *gin.Context)
	ShowDetails(c *gin.Context)
	ShowAddress(c *gin.Context)
	AddAddress(c *gin.Context)
	UpdateAddress(c *gin.Context)
	UpdateUserDetails(c *gin.Context)
	Checkout(c *gin.Context)
	ForgotPassword(c *gin.Context) 
	VerifyResetOtp(c *gin.Context) 
	ResetPassword(c *gin.Context)
}
