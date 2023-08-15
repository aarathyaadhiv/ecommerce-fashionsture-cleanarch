package interfaces

import (
	_ "context"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserSignUp) (models.TokenResponse, error)
	UserLogin(user models.UserLogin) (models.TokenResponse, error)
	// UserOtpLogin(user models.UserOtpLogin)(models.TokenResponse,error)

}
