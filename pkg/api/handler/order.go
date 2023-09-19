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

type OrderHandler struct {
	UseCase services.OrderUseCase
}

func NewOrderHandler(usecase services.OrderUseCase) handler.OrderHandler {
	return &OrderHandler{usecase}
}

// @Summary Place Order
// @Description Place order
// @Tags Order
// @Accept json
// @Produce json
// @Param  OrderRequest body models.OrderRequest true "orderRequest"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders [post]
func (or *OrderHandler) PlaceOrder(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var OrderRequest models.OrderRequest
	if err := c.ShouldBindJSON(&OrderRequest); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := or.UseCase.PlaceOrder(OrderRequest.AddressId, OrderRequest.PaymentId, id.(uint), OrderRequest.CouponId)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully placed order", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Order History
// @Description Showing Order History To User
// @Tags Order
// @Accept json
// @Produce json
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders [get]
func (or *OrderHandler) ShowOrderHistory(c *gin.Context) {
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
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "bad request", nil, errors.New("userid not recovered").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	order, err := or.UseCase.ShowOrderDetails(id.(uint), page, count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing order details", order, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Cancel Order
// @Description Cancel Order By User
// @Tags Order
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders/cancel/{id} [patch]
func (or *OrderHandler) CancelOrder(c *gin.Context) {
	id := c.Param("id")
	err := or.UseCase.CancelOrder(id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully cancelled order", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Approve Order
// @Description Approve Order By Admin
// @Tags Order Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/approval/{id} [patch]
func (or *OrderHandler) AdminApproval(c *gin.Context) {
	id := c.Param("id")
	err := or.UseCase.AdminApproval(id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully approved order", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Return Order
// @Description Return Order
// @Tags Order
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /orders/return/{id} [patch]
func (or *OrderHandler) ReturnOrder(c *gin.Context) {
	id := c.Param("id")
	err := or.UseCase.ReturnOrder(id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully returned order", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show All Orders To Admin
// @Description Show All Orders To Admin
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Param  approval query string true "approval"
// @Param  payment query string true "payment_status"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders [get]
func (or *OrderHandler) ShowOrdersToAdmin(c *gin.Context) {
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
	approval:=c.Query("approval")
	paymentStatus:=c.Query("payment")
	orderDetails, err := or.UseCase.ShowOrderToAdmin(page, count,approval,paymentStatus)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing order details", orderDetails, nil)
	c.JSON(http.StatusOK, succRes)

}

// @Summary Search Order By Admin
// @Description Search Order By Admin
// @Tags Order Management
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  id path string true "id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/orders/{id} [get]
func (or *OrderHandler) SearchOrder(c *gin.Context) {
	id := c.Param("id")
	orderDetails, err := or.UseCase.SearchOrder(id)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing the searched order", orderDetails, nil)
	c.JSON(http.StatusOK, succRes)
}


// @Summary Wallet Of The User
// @Description Wallet Of The User
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /wallet [get]
func (or *OrderHandler) GetWallet(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "error in recovering userid", nil, errors.New("error in fetching userid").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	wallet, err := or.UseCase.GetWallet(id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing wallet", wallet, nil)
	c.JSON(http.StatusOK, succRes)
}
