package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"

type OtpRepository interface {
	CheckUserByPhone(phone string) bool
	FindByPhone(phone string) (models.UserDetails, error)
}
