package interfaces

import "github.com/gin-gonic/gin"


type PaymentHandler interface{
	MakePaymentUsingRazorPay(c *gin.Context)
	VerifyPayment(c *gin.Context)
}