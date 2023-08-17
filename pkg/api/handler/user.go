package handler

import (
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
}

func NewUserHandler(usecase services.UserUseCase) handler.UserHandler {
	return &UserHandler{
		userUseCase: usecase,
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
// @Router /SignUp [post]
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
