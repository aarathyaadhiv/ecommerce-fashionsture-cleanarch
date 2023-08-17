package handler

import (
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
// @Router /admin/adminSignUp [post]
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
// @Router /admin/adminLogin [post]
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
// @Router /admin/user/blockUser/{id} [post]
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
// @Router /admin/user/unblockUser/{id} [post]
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
