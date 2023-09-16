package models

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"

type UserSignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
	PhNo     string `json:"phno" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}
type UserOtpLogin struct {
	PhNo     uint
	Password string
}
type UserDetails struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	PhNo  string `json:"phno"`
}
type TokenResponse struct {
	Token       string
	UserDetails UserDetails
}
type UserLoginCheck struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	PhNo     string `json:"phno"`
	Password string `json:"password"`
}

type ShowAddress struct {
	HouseName string `json:"house_name" binding:"required"`
	Name      string `json:"name" binding:"required"`
	City      string `json:"city" binding:"required"`
	State     string `json:"state" binding:"required"`
	Landmark  string `json:"landmark" binding:"required"`
	Pincode   uint   `json:"pincode" binding:"required"`
}
type UserUpdate struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required" validate:"email"`
	PhNo  string `json:"phno" binding:"required"`
}

type Checkout struct {
	Address       []ShowAddress        `json:"address"`
	PaymentMethod []domain.PaymentMethod `json:"payment_methods"`
	Amount        float64              `json:"amount"`
	Products      []CartProducts       `json:"products"`
}

type Forgot struct {
	Email string `json:"email"  binding:"required" validate:"email"`
}

type ForgotVerify struct {
	Email string `json:"email"  binding:"required" validate:"email"`
	Code  string `json:"code" binding:"required"`
}

type Reset struct {
	NewPassword string `json:"new_password" binding:"required"`
}
