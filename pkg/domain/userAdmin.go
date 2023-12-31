package domain

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"validate:required" `
	Email    string `json:"email" gorm:"validate:required,email" `
	PhNo     string `json:"mobile_number" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
	Role     string `json:"role"`
	Block    bool   `json:"block" gorm:"default:false"`
}

type Address struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	HouseName string `json:"house_name" `
	Name      string `json:"name" `
	City      string `json:"city" gorm:"validate:required"`
	State     string `json:"state" gorm:"validate:required"`
	Landmark  string `json:"landmark" gorm:"validate:required"`
	Pincode   uint   `json:"pincode" gorm:"validate:required"`
	UsersID   uint   `json:"users_id" `
	Users     Users  `json:"users" gorm:"foreignKey:UsersID;constraint:OnDelete:CASCADE"`
}
