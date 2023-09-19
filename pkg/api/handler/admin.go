package handler

import (
	"fmt"
	"net/http"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	AdminUsecase services.AdminUseCase
}

func NewAdminHandler(useCase services.AdminUseCase) handler.AdminHandler {
	return &AdminHandler{
		AdminUsecase: useCase,
	}
}

// @Summary Admin SignUp
// @Description SignUp handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param  admin body models.AdminSignUp true "Admin signup details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/signup [post]
func (ad *AdminHandler) AdminSignUpHandler(c *gin.Context) {
	var adminSignUpDetails models.AdminSignUp

	if err := c.ShouldBindJSON(&adminSignUpDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(adminSignUpDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	adminDetails, err := ad.AdminUsecase.SignUp(adminSignUpDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	sucRes := response.Responses(http.StatusCreated, "successfully signed in", adminDetails, nil)
	c.JSON(http.StatusOK, sucRes)
}

// @Summary Admin Login
// @Description Login handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param  admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/login [post]
func (ad *AdminHandler) AdminLoginHandler(c *gin.Context) {
	var adminLoginDetails models.AdminLogin

	if err := c.ShouldBindJSON(&adminLoginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in the correct order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(adminLoginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraint not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	adminDetails, err := ad.AdminUsecase.Login(adminLoginDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	sucRes := response.Responses(http.StatusOK, "successfully logged in", adminDetails, nil)
	c.JSON(http.StatusOK, sucRes)
}

// @Summary Block User
// @Description Block User By Admin
// @Tags User Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/user/block/{id} [patch]
func (ad *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Param("id")

	if err := ad.AdminUsecase.BlockUser(id); err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully blocked user", nil, nil)
	c.JSON(http.StatusOK, succRes)

}

// @Summary Unblock User
// @Description Unblock User By Admin
// @Tags User Management
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/user/unblock/{id} [patch]
func (ad *AdminHandler) UnblockUser(c *gin.Context) {
	id := c.Param("id")

	if err := ad.AdminUsecase.UnblockUser(id); err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully unblocked user", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary List Users
// @Description List Users To Admin
// @Tags User Management
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  page query string true "page"
// @Param  count query string true "count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/user [get]
func (ad *AdminHandler) ListUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	count := c.DefaultQuery("count", "3")
	users, err := ad.AdminUsecase.ListUsers(page, count)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing users", users, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Admin Home
// @Description Show Details Of Admin
// @Tags Admin Home
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin [get]
func (ad *AdminHandler) AdminHome(c *gin.Context) {
	id, ok := c.Get("adminId")
	if !ok {
		errRes := response.Responses(http.StatusNotFound, "id not recovered from context", nil, fmt.Errorf("id not recovered from context").Error())
		c.JSON(http.StatusNotFound, errRes)
		return
	}
	adminDetails, err := ad.AdminUsecase.AdminHome(id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing admin home page", adminDetails, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary  Sales Report
// @Description  Sales Report
// @Tags Admin Home
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  keyword query string true "keyword"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/salesReport [get]
func (ad *AdminHandler) SalesReport(c *gin.Context) {
	keyword := c.Query("keyword")
	salesReport, err := ad.AdminUsecase.SalesReport(keyword)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing sales report", salesReport, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Admin Dashboard
// @Description Dashboard Of Admin
// @Tags Admin Home
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Failure 500 {object} response.Response{}
// @Router /admin/dashboard [get]
func (ad *AdminHandler) Dashboard(c *gin.Context) {
	dashboard, err := ad.AdminUsecase.Dashboard()
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing dashboard", dashboard, nil)
	c.JSON(http.StatusOK, succRes)
}
