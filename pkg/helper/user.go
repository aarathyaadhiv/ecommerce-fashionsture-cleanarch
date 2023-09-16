package helper

import (
	"regexp"
	"time"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/golang-jwt/jwt"
)

type CustomUserClaim struct {
	ID    uint
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateUserToken(user models.UserDetails) (string, error) {
	claims := &CustomUserClaim{
		ID:    user.ID,
		Email: user.Email,
		Role: "user",
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

func GenerateResetToken(user models.UserLoginCheck) (string, error) {
	claims := &CustomUserClaim{
		ID:    user.ID,
		Email: user.Email,
		Role: "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("reset"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsValidEmail(email string) bool {
    // Define a regex pattern for a valid email address
    pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    match, _ := regexp.MatchString(pattern,email)
    return match
}

func IsValidPhoneNumber(phoneNumber string) bool {
    // Define a regex pattern for the format "+919847256365" (plus sign and 12 digits)
    pattern := `^\+\d{12}$`
    match, _ := regexp.MatchString(pattern, phoneNumber)
    return match
}