package handler

import (
	"errors"
	"fmt"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/copier"

	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	cartUseCase services.CartUseCase
}

func NewUserHandler(usecase services.UserUseCase, cart services.CartUseCase) handler.UserHandler {
	return &UserHandler{
		userUseCase: usecase,
		cartUseCase: cart,
	}
}

// @Summary User SignUp
// @Description SignUp handler for user
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param  user body models.UserSignUp true "user signup details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /signup [post]
func (ur *UserHandler) SignUpHandler(c *gin.Context) {
	var signUpDetails models.UserSignUp
	if err := c.ShouldBindJSON(&signUpDetails); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "fields are not provided in correct format", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}
	if err := validator.New().Struct(signUpDetails); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "constraints not satisfied", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}

	userData, err := ur.userUseCase.UserSignUp(signUpDetails)
	if err != nil {
		erRes := response.Response{Statuscode: http.StatusInternalServerError, Message: "internal server error", Data: nil, Error: err.Error()}
		c.JSON(http.StatusInternalServerError, erRes)
		return
	}
	sucRes := response.Responses(http.StatusCreated, "successfully signedup", userData, nil)
	c.JSON(http.StatusCreated, sucRes)
}

// @Summary User Login
// @Description Login handler for user
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param  user body models.UserLogin true "user login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /login [post]
func (ur *UserHandler) LoginHandler(c *gin.Context) {
	var loginDetails models.UserLogin
	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are  provided in the bad format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(loginDetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userData, err := ur.userUseCase.UserLogin(loginDetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	sucRes := response.Responses(http.StatusOK, "successfully logged in", userData, nil)
	c.JSON(http.StatusOK, sucRes)
}

// @Summary View Details
// @Description View Details Of The User
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /userProfile [get]
func (ur *UserHandler) ShowDetails(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusNotFound, "id not recovered", nil, fmt.Errorf("error in retrieving id from context").Error())
		c.JSON(http.StatusNotFound, errRes)
		return
	}
	userDetails, err := ur.userUseCase.ShowDetails(id.(uint))

	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing user details", userDetails, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Show Address
// @Description Show Addresses of the user
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address [get]
func (ur *UserHandler) ShowAddress(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "id not recovered", nil, fmt.Errorf("error in retrieving id from context").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	address, err := ur.userUseCase.ShowAddress(id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing address", address, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Add Address
// @Description Add Address of the user
// @Tags User Profile
// @Accept json
// @Produce json
// @Param  address body models.ShowAddress true "address"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address [post]
func (ur *UserHandler) AddAddress(c *gin.Context) {
	var address models.ShowAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields is not the rquired order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "error in recovering userid", nil, errors.New("error in fetching userid").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := ur.userUseCase.AddAddress(address, id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully added address", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update Address
// @Description Update Address of the user
// @Tags User Profile
// @Accept json
// @Produce json
// @Param  id path string true "id"
// @Param address body models.ShowAddress true "address"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /address/{id} [patch]
func (ur *UserHandler) UpdateAddress(c *gin.Context) {
	var address models.ShowAddress
	if err := c.ShouldBindJSON(&address); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields is not the rquired order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "error in recovering userid", nil, errors.New("error in fetching userid").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	AddressId := c.Param("id")
	err := ur.userUseCase.UpdateAddress(address, AddressId, id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully updated address", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Update User Details
// @Description Update Details of the user
// @Tags User Profile
// @Accept json
// @Produce json
// @Param  userdetails body models.UserUpdate true "userdetails"
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /userProfile [patch]
func (ur *UserHandler) UpdateUserDetails(c *gin.Context) {
	var userdetails models.UserUpdate
	if err := c.ShouldBindJSON(&userdetails); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields is not the rquired order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "error in recovering userid", nil, errors.New("error in fetching userid").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err := ur.userUseCase.UpdateUserDetails(id.(uint), userdetails)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server errorr", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully updated user details", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary CheckOut
// @Description Displaying Checkout Page Before Placing Order
// @Tags Checkout Page
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /checkout [get]
func (ur *UserHandler) Checkout(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		errRes := response.Responses(http.StatusBadRequest, "error in recovering userid", nil, errors.New("error in fetching userid").Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	checkoutDetails, err := ur.userUseCase.Checkout(id.(uint))
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully showing checkout page", checkoutDetails, nil)
	c.JSON(http.StatusOK, succRes)

}

// @Summary forgot password
// @Description forgot password for user to reset password by sending otp to phone number
// @Tags User Profile
// @Accept json
// @Produce json
// @Param  forgot body models.Forgot true "forgot"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgotPassword [post]
func (ur *UserHandler) ForgotPassword(c *gin.Context) {
	var forgot models.Forgot
	if err := c.ShouldBindJSON(&forgot); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields is not the rquired order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(forgot); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "constraints not satisfied", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}
	err := ur.userUseCase.ForgotPassword(forgot)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully sent otp", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Verify Reset Otp
// @Description Verify otp to reset password
// @Tags User Profile
// @Accept json
// @Produce json
// @Param  verify body models.ForgotVerify true "verify"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgotPassword/verify [post]
func (ur *UserHandler) VerifyResetOtp(c *gin.Context) {
	var verify models.ForgotVerify

	if err := c.ShouldBindJSON(&verify); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	tokenString, err := ur.userUseCase.VerifyResetOtp(verify)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "otp verified", tokenString, nil)
	c.JSON(http.StatusOK, succRes)
}

// @Summary Reset Password
// @Description To Reset Password
// @Tags User Profile
// @Accept json
// @Produce json
// @Security ApiKeyHeaderAuth
// @Param  reset body models.Reset true "reset"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /forgotPassword/reset [post]
func (ur *UserHandler) ResetPassword(c *gin.Context) {
	var reset models.Reset
	if err := c.ShouldBindJSON(&reset); err != nil {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "fields are not provided in correct format", Data: nil, Error: err.Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}
	id, ok := c.Get("resetId")
	if !ok {
		erRes := response.Response{Statuscode: http.StatusBadRequest, Message: "reset id not retrieved", Data: nil, Error: errors.New("error in retrieving reset id").Error()}
		c.JSON(http.StatusBadRequest, erRes)
		return
	}
	err := ur.userUseCase.ResetPassword(id.(uint), reset.NewPassword)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully reset password", nil, nil)
	c.JSON(http.StatusOK, succRes)
}
