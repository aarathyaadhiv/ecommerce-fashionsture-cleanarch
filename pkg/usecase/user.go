package usecase

import (
	_ "context"
	"errors"
	"strconv"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
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
	cartRepo repo.CartRepository
	config   config.Config
}

func NewUserUseCase(repo repo.UserRepository, cart repo.CartRepository, config config.Config) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
		cartRepo: cart,
		config:   config,
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

func (c *userUseCase) ShowDetails(id uint) (models.UserDetails, error) {
	return c.userRepo.FindByID(id)
}

func (c *userUseCase) ShowAddress(id uint) ([]models.ShowAddress, error) {
	return c.userRepo.ShowAddress(id)
}

func (c *userUseCase) AddAddress(address models.ShowAddress, userId uint) error {
	return c.userRepo.AddAddress(address, userId)
}

func (c *userUseCase) UpdateAddress(address models.ShowAddress, addressId string, userId uint) error {
	id, err := strconv.Atoi(addressId)
	if err != nil {
		return err
	}
	return c.userRepo.UpdateAddress(address, uint(id), userId)
}

func (c *userUseCase) UpdateUserDetails(userId uint, userdetails models.UserUpdate) error {
	return c.userRepo.UpdateUserDetails(userId, userdetails)
}

func (c *userUseCase) Checkout(id uint) (models.Checkout, error) {
	address, err := c.userRepo.ShowAddress(id)
	if err != nil {
		return models.Checkout{}, err
	}

	products, err := c.cartRepo.ShowProductInCart(id)
	if err != nil {
		return models.Checkout{}, err
	}

	amount, err := c.cartRepo.TotalAmountInCart(id)
	if err != nil {
		return models.Checkout{}, err
	}

	payment, err := c.cartRepo.PaymentMethods()
	if err != nil {
		return models.Checkout{}, err
	}

	return models.Checkout{
		Address:       address,
		Amount:        amount,
		Products:      products,
		PaymentMethod: payment,
	}, nil
}

func (c *userUseCase) ForgotPassword(forgot models.Forgot) error {
	if ok := c.userRepo.CheckUserAvailability(forgot.Email); !ok {
		return errors.New("no such user exist")
	}
	if ok := c.userRepo.IsBlocked(forgot.Email); ok {
		return errors.New("user is blocked")
	}
	user, err := c.userRepo.FindByEmail(forgot.Email)
	if err != nil {
		return err
	}
	helper.TwilioSetUp(c.config.TwilioAccountSID, c.config.TwilioAuthToken)
	_, err = helper.TwilioSendOTP(user.PhNo, c.config.TwilioServicesId)
	return err

}

func (c *userUseCase) VerifyResetOtp(data models.ForgotVerify) (string, error) {
	if ok := c.userRepo.CheckUserAvailability(data.Email); !ok {
		return "", errors.New("no such user exist")
	}
	if ok := c.userRepo.IsBlocked(data.Email); ok {
		return "", errors.New("user is blocked")
	}
	user, err := c.userRepo.FindByEmail(data.Email)
	if err != nil {
		return "", err
	}
	verify := models.VerifyData{PhoneNumber: user.PhNo, Code: data.Code}
	helper.TwilioSetUp(c.config.TwilioAccountSID, c.config.TwilioAuthToken)
	if err := helper.TwilioVerifyOTP(verify, c.config.TwilioServicesId); err != nil {
		return "", err
	}

	tokenString, err := helper.GenerateResetToken(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (c *userUseCase) ResetPassword(id uint, password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	return c.userRepo.UpdatePassword(id, string(hashPassword))
}
