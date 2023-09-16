package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"

type AdminUseCase interface {
	SignUp(admin models.AdminSignUp) (models.AdminTokenResponse, error)
	Login(admin models.AdminLogin) (models.AdminTokenResponse, error)
	BlockUser(id string) error
	UnblockUser(id string) error
	ListUsers(pages,counts string)([]models.AdminUserResponse,error)
	AdminHome(id uint)(models.AdminDetails,error)
	Dashboard() (models.Dashboard, error)
	SalesReport(timeWord string)(models.SalesReport,error)
}
