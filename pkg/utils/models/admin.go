package models

type AdminSignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
	PhNo     string `json:"phno" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

type AdminDetails struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	PhNo  string `json:"phno"`
}
type AdminTokenResponse struct {
	Token        string
	AdminDetails AdminDetails
}
type AdminSignUpResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	PhNo     string `json:"phno"`
	Password string `json:"password"`
}
