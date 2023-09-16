package interfaces

import (
	"time"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type AdminRepository interface {
	CheckAdminAvailability(email string) bool
	FindByEmail(email string) (models.AdminSignUpResponse, error)
	Save(admin models.AdminSignUp) (models.AdminDetails, error)
	BlockUser(id uint) error
	UnblockUser(id uint) error
	IsBlocked(id uint) bool
	ListUsers(page,count int)([]models.AdminUserResponse,error)
	AdminDetails(id uint)(models.AdminDetails,error)
	DashboardRevenue()(models.DashboardRevenue,error)
	DashboardOrders()(models.DashboardOrders,error)
	DashboardAmount()(models.DashboardAmount,error)
	DashboardUsers()(models.DashboardUsers,error)
	DashboardProduct()(models.DashboardProduct,error)
	SalesReport(startDate,endDate time.Time)(models.SalesReport,error)
}
