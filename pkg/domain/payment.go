package domain


type PaymentMethod struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Method string `json:"method"`
}