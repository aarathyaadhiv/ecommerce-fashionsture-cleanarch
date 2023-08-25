package interfaces

import (
	_ "context"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserSignUp) (models.TokenResponse, error)
	UserLogin(user models.UserLogin) (models.TokenResponse, error)
	ShowDetails(id uint)(models.UserDetails,error)
	ShowAddress(id uint)([]models.ShowAddress,error)
	AddAddress(address models.ShowAddress,userId uint)error
	UpdateAddress(address models.ShowAddress,addressId string,userId uint)error
	UpdateUserDetails(userId uint,userdetails models.UserUpdate)error
	Checkout(id uint )(models.Checkout,error)
	ForgotPassword(forgot models.Forgot)(error)
	VerifyResetOtp(data models.ForgotVerify) (string, error)
	ResetPassword(id uint,password string)error
}
