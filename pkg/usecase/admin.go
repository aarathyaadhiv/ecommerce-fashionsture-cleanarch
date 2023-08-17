package usecase

import (
	"errors"
	"strconv"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo repo.AdminRepository
}

func NewAdminUseCase(repo repo.AdminRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo: repo,
	}
}

func (c *AdminUseCase) SignUp(admin models.AdminSignUp) (models.AdminTokenResponse, error) {
	if ok := c.adminRepo.CheckAdminAvailability(admin.Email); ok {
		return models.AdminTokenResponse{}, errors.New("already existing email")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in password hashing")
	}
	admin.Password = string(hashPassword)

	adminDetails, err := c.adminRepo.Save(admin)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in saving user data")
	}

	tokenString, err := helper.GenerateAdminToken(adminDetails)
	if err != nil {
		return models.AdminTokenResponse{}, err
	}
	return models.AdminTokenResponse{AdminDetails: adminDetails, Token: tokenString}, nil
}

func (c *AdminUseCase) Login(admin models.AdminLogin) (models.AdminTokenResponse, error) {
	if ok := c.adminRepo.CheckAdminAvailability(admin.Email); !ok {
		return models.AdminTokenResponse{}, errors.New("no such user exist")
	}

	adminCompare, err := c.adminRepo.FindByEmail(admin.Email)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(adminCompare.Password), []byte(admin.Password)); err != nil {
		return models.AdminTokenResponse{}, errors.New("password is incorrect")
	}

	var adminDetails models.AdminDetails
	if err := copier.Copy(&adminDetails, &adminCompare); err != nil {
		return models.AdminTokenResponse{}, err
	}
	tokenString, err := helper.GenerateAdminToken(adminDetails)
	if err != nil {
		return models.AdminTokenResponse{}, err
	}

	return models.AdminTokenResponse{AdminDetails: adminDetails, Token: tokenString}, nil
}

func (c *AdminUseCase) BlockUser(id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if ok:=c.adminRepo.IsBlocked(uint(userId));ok{
		return errors.New("already blocked user")
	}
	return c.adminRepo.BlockUser(uint(userId))
}

func (c *AdminUseCase) UnblockUser(id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if ok:=c.adminRepo.IsBlocked(uint(userId));!ok{
		return errors.New("already unblocked user")
	}
	return c.adminRepo.UnblockUser(uint(userId))
}
