package repository

import (
	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type OtpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &OtpRepository{DB}
}

func (c *OtpRepository) CheckUserByPhone(phone string) bool {
	var count int
	if err := c.DB.Raw(`select count(*) from users where role='user' and ph_no=?`, phone).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *OtpRepository) FindByPhone(phone string) (models.UserDetails, error) {
	var userDetails models.UserDetails
	if err := c.DB.Raw(`select id,name,email,ph_no from users where ph_no=?`, phone).Scan(&userDetails).Error; err != nil {
		return models.UserDetails{}, err
	}
	return userDetails, nil
}
