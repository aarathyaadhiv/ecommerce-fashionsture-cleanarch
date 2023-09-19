package handler

import (
	"errors"
	"net/http"
	"strconv"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	_"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Usecase usecase.CartUseCase
}

func NewCartHandler(usecase usecase.CartUseCase) handler.CartHandler {
	return &CartHandler{usecase}
}
// @Summary Add To Cart
// @Description Add Products To Cart
// @Tags Cart 
// @Accept json
// @Produce json
// @Param  id path string true "product_id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/{id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productId:=c.Param("id")
	product,err:=strconv.Atoi(productId)
	if err!=nil{
		errRes:=response.Responses(http.StatusBadRequest,"query param recovery error",nil,errors.New("error in fetching path params").Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	
	
	err=cr.Usecase.AddToCart(id.(uint),uint(product))
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfuuly added to cart",nil,nil)
	c.JSON(http.StatusOK,succRes)
	
}
// @Summary Remove From Cart 
// @Description Remove Product From Cart
// @Tags Cart 
// @Accept json
// @Produce json
// @Param  id path string true "product_id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart/remove/{id} [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context){
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productId:=c.Param("id")

	err:=cr.Usecase.RemoveFromCart(id.(uint),productId)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	
	succRes:=response.Responses(http.StatusOK,"successfuuly removed from cart",nil,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Show Cart Products 
// @Description Show Products In Users Cart
// @Tags Cart 
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart [get]
func (cr *CartHandler) ShowProductInCart(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	products,err:=cr.Usecase.ShowProductInCart(id.(uint),page,count)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully showing products in cart",products,nil)
	c.JSON(http.StatusOK,succRes)
}

// @Summary empty Cart Products 
// @Description empty Products In Users Cart
// @Tags Cart 
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /cart [delete]
func(cr *CartHandler) EmptyCart(c *gin.Context){
	id,ok:=c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "userid not retrieved", nil, errors.New("userid retrieval error").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err:=cr.Usecase.EmptyCart(id.(uint))
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successfully empty cart",nil,nil)
	c.JSON(http.StatusOK,succRes)
}