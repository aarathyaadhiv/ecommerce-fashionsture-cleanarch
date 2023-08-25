package handler

import (
	"errors"
	"net/http"
	"strconv"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)



type OrderHandler struct{
	UseCase services.OrderUseCase
}

func NewOrderHandler(usecase services.OrderUseCase)handler.OrderHandler{
	return &OrderHandler{usecase}
}
// @Summary Place Order
// @Description Place order
// @Tags Order Management
// @Accept json
// @Produce json
// @Param  OrderRequest body models.OrderRequest true "orderRequest"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders/placeOrder [post]
func (or *OrderHandler) PlaceOrder(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.Responses(http.StatusBadRequest,"bad request",nil,errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	var OrderRequest models.OrderRequest
	if err:=c.ShouldBindJSON(&OrderRequest);err!=nil{
		errRes:=response.Responses(http.StatusBadRequest,"bad request",nil,errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	err:=or.UseCase.PlaceOrder(OrderRequest.AddressId,OrderRequest.PaymentId,id.(uint))
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully placed order",nil,nil)
	c.JSON(http.StatusOK,succRes)
}

// @Summary Order History
// @Description Showing Order History To User
// @Tags Order Management
// @Accept json
// @Produce json
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders [get]
func (or *OrderHandler) ShowOrderHistory(c *gin.Context){
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "4"))
	if err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id,ok:=c.Get("userId")
	if !ok{
		errRes:=response.Responses(http.StatusBadRequest,"bad request",nil,errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	order,err:=or.UseCase.ShowOrderDetails(id.(uint),page,count)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully showing order details",order,nil)
	c.JSON(http.StatusOK,succRes)
}