package repository

import (
	"errors"
	"time"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) repo.AdminRepository {
	return &AdminRepository{DB}
}

func (c *AdminRepository) CheckAdminAvailability(email string) bool {
	var count int
	if err := c.DB.Raw(`select count(*) from users where email=? and role='admin'`, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *AdminRepository) FindByEmail(email string) (models.AdminSignUpResponse, error) {
	var adminDetails models.AdminSignUpResponse
	if err := c.DB.Raw(`select id,name,email,ph_no,password from users where email=? and role='admin'`, email).Scan(&adminDetails).Error; err != nil {
		return models.AdminSignUpResponse{}, errors.New("error in fetching admin details")
	}
	return adminDetails, nil
}
func (c *AdminRepository) Save(admin models.AdminSignUp) (models.AdminDetails, error) {
	var adminDetails models.AdminDetails
	if err := c.DB.Raw(`insert into users(name,email,ph_no,password,role) values($1,$2,$3,$4,$5) returning id,name,email,ph_no`, admin.Name, admin.Email, admin.PhNo, admin.Password, "admin").Scan(&adminDetails).Error; err != nil {
		return models.AdminDetails{}, errors.New("error in fetching admin details")
	}
	return adminDetails, nil
}

func (c *AdminRepository) BlockUser(id uint) error {
	return c.DB.Exec(`update users set block='true' where id=?`, id).Error
}

func (c *AdminRepository) UnblockUser(id uint) error {
	return c.DB.Exec(`update users set block='false' where id=?`, id).Error
}

func (c *AdminRepository) IsBlocked(id uint) bool {
	var block bool
	if err := c.DB.Raw(`select block from users where id=?`, id).Scan(&block).Error; err != nil {
		return false
	}
	return block
}

func (c *AdminRepository) ListUsers(page,count int) ([]models.AdminUserResponse, error) {
	offset:=(page-1)*count
	var users []models.AdminUserResponse
	err := c.DB.Raw(`SELECT id,name,email,ph_no,block as status FROM users WHERE role='user' limit ? offset ?`,count,offset).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *AdminRepository) AdminDetails(id uint) (models.AdminDetails, error) {
	var adminDetails models.AdminDetails
	err := c.DB.Raw(`SELECT id,name,email,ph_no FROM users WHERE role='admin'`).Scan(&adminDetails).Error
	if err != nil {
		return models.AdminDetails{}, err
	}
	return adminDetails, nil
}

func (c *AdminRepository) DashboardRevenue() (models.DashboardRevenue, error) {
	var dashboard models.DashboardRevenue
	err := c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='paid' AND order_date<? AND order_date>?`, time.Now(), time.Now().AddDate(0, 0, -1)).Scan(&dashboard.DayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, err
	}
	err = c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='paid' AND order_date<? AND order_date>?`, time.Now(), time.Now().AddDate(0, -1, 0)).Scan(&dashboard.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, err
	}
	err = c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='paid' AND order_date<? AND order_date>?`, time.Now(), time.Now().AddDate(-1, 0, 0)).Scan(&dashboard.YearlyRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, err
	}
	return dashboard, nil
}

func (c *AdminRepository) DashboardOrders() (models.DashboardOrders, error) {
	var dashboard models.DashboardOrders
	err := c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE payment_status='paid'`).Scan(&dashboard.CompleteOrder).Error
	if err != nil {
		return models.DashboardOrders{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE payment_status='not paid'`).Scan(&dashboard.PendingOrder).Error
	if err != nil {
		return models.DashboardOrders{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE status='cancelled'`).Scan(&dashboard.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrders{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders `).Scan(&dashboard.TotalOrder).Error
	if err != nil {
		return models.DashboardOrders{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(DISTINCT users_id) FROM orders `).Scan(&dashboard.TotalOrderedUsers).Error
	if err != nil {
		return models.DashboardOrders{}, err
	}
	return dashboard, nil
}

func (c *AdminRepository) DashboardAmount() (models.DashboardAmount, error) {
	var dashboard models.DashboardAmount
	err := c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='paid'`).Scan(&dashboard.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, err
	}
	err = c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='not paid'`).Scan(&dashboard.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, err
	}
	return dashboard, nil
}

func (c *AdminRepository) DashboardUsers() (models.DashboardUsers, error) {
	var dashboard models.DashboardUsers
	err := c.DB.Raw(`SELECT COUNT(*) FROM users WHERE role='user'`).Scan(&dashboard.TotalUsers).Error
	if err != nil {
		return models.DashboardUsers{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM users WHERE role='user' AND block=true`).Scan(&dashboard.BlockedUsers).Error
	if err != nil {
		return models.DashboardUsers{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(DISTINCT users_id) FROM orders `).Scan(&dashboard.OrderedUsers).Error
	if err != nil {
		return models.DashboardUsers{}, err
	}
	return dashboard, nil
}

func (c *AdminRepository) DashboardProduct() (models.DashboardProduct, error) {
	var dashboard models.DashboardProduct
	err := c.DB.Raw(`SELECT COUNT(*) FROM products `).Scan(&dashboard.TotalProducts).Error
	if err != nil {
		return models.DashboardProduct{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM products WHERE quantity=0`).Scan(&dashboard.OutOfStockProducts).Error
	if err != nil {
		return models.DashboardProduct{}, err
	}
	var productId uint
	err = c.DB.Raw(`SELECT product_id FROM order_products GROUP BY product_id ORDER BY MAX(quantity) DESC LIMIT 1`).Scan(&productId).Error
	if err != nil {
		return models.DashboardProduct{}, err
	}
	err = c.DB.Raw(`SELECT name FROM products WHERE id=?`, productId).Scan(&dashboard.TopSellingProduct).Error
	if err != nil {
		return models.DashboardProduct{}, err
	}
	return dashboard, nil
}

func (c *AdminRepository) SalesReport(startDate, endDate time.Time) (models.SalesReport, error) {
	var salesReport models.SalesReport
	err := c.DB.Raw(`SELECT coalesce(SUM(amount),0.00) FROM orders WHERE payment_status='paid' AND order_date>? AND order_date<?`, startDate, endDate).Scan(&salesReport.TotalSales).Error
	if err != nil {
		return models.SalesReport{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE order_date>? AND order_date<?`, startDate, endDate).Scan(&salesReport.TotalOrders).Error
	if err != nil {
		return models.SalesReport{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE payment_status='paid' AND order_date>? AND order_date<?`, startDate, endDate).Scan(&salesReport.CompleteOrders).Error
	if err != nil {
		return models.SalesReport{}, err
	}
	err = c.DB.Raw(`SELECT COUNT(*) FROM orders WHERE payment_status='not paid' AND order_date>? AND order_date<?`, startDate, endDate).Scan(&salesReport.PendingOrder).Error
	if err != nil {
		return models.SalesReport{}, err
	}
	return salesReport, nil
}
