package interfaces

import (
	_ "context"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	FindByEmail(email string) (models.UserLoginCheck, error)
	FindAll() ([]models.UserDetails, error)
	FindByID(id uint) (models.UserDetails, error)
	Save(user models.UserSignUp) (models.UserDetails, error)
	IsBlocked(email string) bool
	ShowAddress(id uint) ([]models.ShowAddress, error)
	AddAddress(address models.ShowAddress,userId uint)error
	UpdateAddress(address models.ShowAddress,addressId,userId uint)error
	UpdateUserDetails(userId uint,userDetails models.UserUpdate)error
	UpdatePassword(id uint,password string)error
}
