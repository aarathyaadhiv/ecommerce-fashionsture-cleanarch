package interfaces

import "github.com/gin-gonic/gin"

type OtpHandler interface {
	SendOTP(c *gin.Context)
	VerifyOtp(c *gin.Context)
}
