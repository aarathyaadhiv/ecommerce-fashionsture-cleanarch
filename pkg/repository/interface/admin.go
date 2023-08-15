package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"

type AdminRepository interface {
	CheckAdminAvailability(email string) bool
	FindByEmail(email string) (models.AdminSignUpResponse, error)
	Save(admin models.AdminSignUp) (models.AdminDetails, error)
	BlockUser(id uint) error
	UnblockUser(id uint) error
}
