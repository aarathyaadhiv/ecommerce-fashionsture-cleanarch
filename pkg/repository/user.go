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

	if err := c.DB.Raw(`select id,name,email,ph_no as phno,password from users where email=? and role='user'`, email).Scan(&user).Error; err != nil {
		return models.UserLoginCheck{}, errors.New("error in checking userdetails")
	}
	return user, nil
}

func (c *userDatabase) FindAll(page,count int) ([]models.UserDetails, error) {
	var users []models.UserDetails
	offset:=(page-1)*count
	if err := c.DB.Raw(`select id,name,email,ph_no from users where role='user' limit ? offset ?`,count,offset).Scan(&users).Error; err != nil {
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
		return models.UserDetails{}, errors.New("error saving in database")
	}

	return userdetails, nil
}

func (c *userDatabase) IsBlocked(email string) (bool,error) {
	var block bool
	if err := c.DB.Raw(`select block from users where email=?`, email).Scan(&block).Error; err != nil {
		return false,errors.New("error in fetching block detail")
	}
	return block,nil
}

func (c *userDatabase) ShowAddress(id uint,page,count int) ([]models.ShowAddress, error) {
	var showAddress []models.ShowAddress
	offset:=(page-1)*count
	if err := c.DB.Raw(`SELECT house_name,name,city,state,landmark,pincode FROM addresses WHERE users_id=? limit ? offset ?`, id,count,offset).Scan(&showAddress).Error; err != nil {
		return nil, err
	}
	return showAddress, nil
}

func (c *userDatabase) AddAddress(address models.ShowAddress, userId uint) error {
	return c.DB.Exec(`INSERT INTO addresses(house_name,name,city,state,landmark,pincode,users_id ) VALUES(?,?,?,?,?,?,?)`, address.HouseName, address.Name, address.City, address.State, address.Landmark, address.Pincode, userId).Error
}

func (c *userDatabase) UpdateAddress(address models.ShowAddress, addressId, userId uint) error {
	return c.DB.Exec(`UPDATE addresses SET house_name=?,name=?,city=?,state=?,landmark=?,pincode=? WHERE id=? AND users_id=?`, address.HouseName, address.Name, address.City, address.State, address.Landmark, address.Pincode, addressId, userId).Error
}

func (c *userDatabase) UpdateUserDetails(userId uint, userDetails models.UserUpdate) error {
	err:= c.DB.Exec(`UPDATE users SET name=?,email=?,ph_no=? WHERE id=?`, userDetails.Name, userDetails.Email, userDetails.PhNo, userId).Error
	if err!=nil{
		return errors.New("error while updating")
	}
	return nil
}

func (c *userDatabase) UpdatePassword(id uint, password string) error {
	return c.DB.Exec(`UPDATE users SET password=? WHERE id=?`, password, id).Error
}
