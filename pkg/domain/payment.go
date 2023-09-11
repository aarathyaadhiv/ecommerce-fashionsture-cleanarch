package domain


type PaymentMethod struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Method string `json:"method"`
}

type RazorPay struct{
	ID uint `json:"id" gorm:"primaryKey"`
	RazorId string`json:"razor_id" `
	OrderId uint `json:"order_id"`
	PaymentId string `json:"payment_id"`
}