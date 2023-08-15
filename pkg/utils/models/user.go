package models

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
