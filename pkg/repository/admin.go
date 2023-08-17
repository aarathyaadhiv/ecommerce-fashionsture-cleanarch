package repository

import (
	"errors"

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