package models

type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}
