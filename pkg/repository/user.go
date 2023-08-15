package repository

import (
	_ "context"
	"errors"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) repo.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int

	if err := c.DB.Raw(`select count(*) from users where email=? and role='user'`, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (c *userDatabase) FindByEmail(email string) (models.UserLoginCheck, error) {
	var user models.UserLoginCheck

	if err := c.DB.Raw(`select id,name,email,Ph_no,password from users where email=? and role='user'`, email).Scan(&user).Error; err != nil {
		return models.UserLoginCheck{}, errors.New("error in checking userdetails")
	}
	return user, nil
}

func (c *userDatabase) FindAll() ([]models.UserDetails, error) {
	var users []models.UserDetails
	if err := c.DB.Raw(`select id,name,email,ph_no from users where role='user'`).Scan(&users).Error; err != nil {
		return []models.UserDetails{}, errors.New("error in fetching userdetails")
	}

	return users, nil
}

func (c *userDatabase) FindByID(id uint) (models.UserDetails, error) {
	var user models.UserDetails
	if err := c.DB.Raw(`select id,name,email,ph_no from users where role='user' and id=?`, id).Scan(&user).Error; err != nil {
		return models.UserDetails{}, errors.New("error in fetching userdetails")
	}

	return user, nil
}

func (c *userDatabase) Save(user models.UserSignUp) (models.UserDetails, error) {
	var userdetails models.UserDetails
	if err := c.DB.Raw(`insert into users(name,email,ph_no,password,role) values($1,$2,$3,$4,$5) returning id,name,email,ph_no`, user.Name, user.Email, user.PhNo, user.Password, "user").Scan(&userdetails).Error; err != nil {
		return models.UserDetails{}, errors.New("error in saving in database")
	}

	return userdetails, nil
}

func (c *userDatabase) IsBlocked(email string) bool {
	var block bool
	if err := c.DB.Raw(`select block from users where email=?`, email).Scan(&block).Error; err != nil {
		return false
	}
	return block
}
