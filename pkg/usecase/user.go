package usecase

import (
	_ "context"
	"errors"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignUp(user models.UserSignUp) (models.TokenResponse, error) {
	if ok := c.userRepo.CheckUserAvailability(user.Email); ok {
		return models.TokenResponse{}, errors.New("already existing email")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenResponse{}, errors.New("error in password hashing")
	}
	user.Password = string(hashPassword)

	userDetails, err := c.userRepo.Save(user)
	if err != nil {
		return models.TokenResponse{}, errors.New("error in saving user data")
	}

	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return models.TokenResponse{}, err
	}
	return models.TokenResponse{UserDetails: userDetails, Token: tokenString}, nil
}
func (c *userUseCase) UserLogin(user models.UserLogin) (models.TokenResponse, error) {
	if ok := c.userRepo.CheckUserAvailability(user.Email); !ok {
		return models.TokenResponse{}, errors.New("no such user exist")
	}
	if ok := c.userRepo.IsBlocked(user.Email); ok {
		return models.TokenResponse{}, errors.New("user is blocked")
	}
	userCompare, err := c.userRepo.FindByEmail(user.Email)
	if err != nil {
		return models.TokenResponse{}, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userCompare.Password), []byte(user.Password)); err != nil {
		return models.TokenResponse{}, errors.New("password is incorrect")
	}

	var userDetails models.UserDetails
	if err := copier.Copy(&userDetails, &userCompare); err != nil {
		return models.TokenResponse{}, err
	}
	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return models.TokenResponse{UserDetails: userDetails, Token: tokenString}, nil

}
