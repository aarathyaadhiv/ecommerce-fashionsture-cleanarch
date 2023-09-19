package handler

import (
	"net/http"
	"time"

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
	parsedExpiry,err:=time.Parse("2006-01-02",coupon.Expiry)
	if err!=nil{
		errRes := response.Responses(http.StatusBadRequest, "expiry date is not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	coupon.ExpiryTime=parsedExpiry
	err=co.Usecase.AddCoupon(coupon)
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
// @Summary Show Coupons 
// @Description Show Coupons To Admin
// @Tags Coupon Management
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupon [get]
func (co *CouponHandler) GetCoupon(c *gin.Context){
	page:=c.DefaultQuery("page","1")
	count:=c.DefaultQuery("count","3")
	coupon,err:=co.Usecase.GetCoupon(page,count)
	if err!=nil{
		errRes:=response.Responses(http.StatusInternalServerError,"internal server error",nil,err.Error())
		c.JSON(http.StatusInternalServerError,errRes)
		return
	}
	succRes:=response.Responses(http.StatusOK,"successsfully showing coupons",coupon,nil)
	c.JSON(http.StatusOK,succRes)
}