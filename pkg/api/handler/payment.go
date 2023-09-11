package handler

import (
	"fmt"
	"net/http"

	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)



type PaymentHandler struct{
	Usecase usecase.PaymentUseCase
}

func NewPaymentHandler(usecase usecase.PaymentUseCase)interfaces.PaymentHandler{
	return &PaymentHandler{usecase}
}


func (ph *PaymentHandler) MakePaymentUsingRazorPay(c *gin.Context){
	id:=c.Param("order_id")
	order,err:=ph.Usecase.RazorPayPayment(id)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}

	c.HTML(http.StatusOK,"index.html",order)

}

func (ph *PaymentHandler) VerifyPayment(c *gin.Context){
	orderId:=c.Query("order_id")
	paymentId:=c.Query("payment_id")
	razorId:=c.Query("razor_id")
	fmt.Println("code enters")
	fmt.Println(orderId)
	fmt.Println(razorId)
	fmt.Println(paymentId)
	err:=ph.Usecase.SaveRazorPayPaymentId(orderId,razorId,paymentId)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully saved payment details",nil,nil)
	c.JSON(http.StatusOK,succRes)

}