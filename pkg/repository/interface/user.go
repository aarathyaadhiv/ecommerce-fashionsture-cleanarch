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
}
