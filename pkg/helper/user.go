package helper

import (
	"time"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/golang-jwt/jwt"
)

type CustomUserClaim struct {
	ID    uint
	Email string
	jwt.StandardClaims
}

func GenerateUserToken(user models.UserDetails) (string, error) {
	claims := &CustomUserClaim{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
