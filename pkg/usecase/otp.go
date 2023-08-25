package usecase

import (
	"errors"
	

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type OtpUseCase struct {
	config config.Config
	repo   repo.OtpRepository
}

func NewOtpUseCase(config config.Config, repo repo.OtpRepository) usecase.OtpUseCase {
	return &OtpUseCase{
		config: config,
		repo:   repo,
	}
}
func (c *OtpUseCase) SendOTP(user models.OTPData) error {
	if ok := c.repo.CheckUserByPhone(user.PhoneNumber); !ok {
		return errors.New("no user exist with this phone no")
	}
	helper.TwilioSetUp(c.config.TwilioAccountSID, c.config.TwilioAuthToken)
	_, err := helper.TwilioSendOTP(user.PhoneNumber, c.config.TwilioServicesId)
	if err != nil {
		
		return err
	}
	return nil

}

func (c *OtpUseCase) VerifyOTP(data models.VerifyData) (models.TokenResponse, error) {
	helper.TwilioSetUp(c.config.TwilioAccountSID, c.config.TwilioAuthToken)
	if err := helper.TwilioVerifyOTP(data, c.config.TwilioServicesId); err != nil {
		return models.TokenResponse{}, err
	}
	userDetails, err := c.repo.FindByPhone(data.PhoneNumber)
	if err != nil {
		return models.TokenResponse{}, err
	}
	tokenString, err := helper.GenerateUserToken(userDetails)
	if err != nil {
		return models.TokenResponse{}, err
	}

	return models.TokenResponse{UserDetails: userDetails, Token: tokenString}, nil
}



