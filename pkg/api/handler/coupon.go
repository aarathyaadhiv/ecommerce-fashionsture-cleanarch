package handler

import (
	"net/http"

	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)


type CouponHandler struct{
	Usecase usecase.CouponUseCase
}


func NewCouponHandler(usecase usecase.CouponUseCase)interfaces.CouponHandler{
	return &CouponHandler{usecase}
}

// @Summary Add coupon
// @Description Add coupon 
// @Tags Coupon Management
// @Accept json
// @Produce json
// @Param  coupon body models.AddCoupon true "coupon details"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon [post]
func (co *CouponHandler) AddCoupon(c *gin.Context){
	var coupon models.AddCoupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err:=co.Usecase.AddCoupon(coupon.CouponId,coupon.Discount,coupon.Usage)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successsfully added coupon",nil,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Expire Coupon 
// @Description Expire Coupon By Admin
// @Tags Coupon Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/expire/{id} [patch]
func (co *CouponHandler) ExpireCoupon(c *gin.Context){
	id:=c.Param("id")
	err:=co.Usecase.ExpireCoupon(id)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successsfully expired coupon",nil,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Block Coupon 
// @Description Block Coupon By Admin
// @Tags Coupon Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/block/{id} [patch]
func (co *CouponHandler) BlockCoupon(c *gin.Context){
	id:=c.Param("id")
	err:=co.Usecase.BlockCoupon(id)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successsfully blocked coupon",nil,nil)
	c.JSON(http.StatusOK,succRes)
}
// @Summary Unblock Coupon 
// @Description Unblock Coupon By Admin
// @Tags Coupon Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon/unblock/{id} [patch]
func (co *CouponHandler) UnBlockCoupon(c *gin.Context){
	id:=c.Param("id")
	err:=co.Usecase.UnBlockCoupon(id)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successsfully unblocked coupon",nil,nil)
	c.JSON(http.StatusOK,succRes)
}