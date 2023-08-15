package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"

type OtpUseCase interface {
	SendOTP(user models.OTPData) error
	VerifyOTP(data models.VerifyData) (models.TokenResponse, error)
}
