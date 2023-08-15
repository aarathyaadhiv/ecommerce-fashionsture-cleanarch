package handler

import (
	"net/http"

	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler/interface"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	UseCase usecase.OtpUseCase
}

func NewOtpHandler(usecase usecase.OtpUseCase) interfaces.OtpHandler {
	return &OtpHandler{usecase}
}

func (ot *OtpHandler) SendOTP(c *gin.Context) {
	var otp models.OTPData
	if err := c.ShouldBindJSON(&otp); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := ot.UseCase.SendOTP(otp); err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully sent otp", nil, nil)
	c.JSON(http.StatusOK, succRes)
}

func (ot *OtpHandler) VerifyOtp(c *gin.Context) {
	var verify models.VerifyData

	if err := c.ShouldBindJSON(&verify); err != nil {
		errRes := response.Responses(http.StatusBadRequest, "fields are not in correct manner", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	userDetails, err := ot.UseCase.VerifyOTP(verify)
	if err != nil {
		errRes := response.Responses(http.StatusInternalServerError, "internal server error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	succRes := response.Responses(http.StatusOK, "successfully logged in", userDetails, nil)
	c.JSON(http.StatusOK, succRes)
}
