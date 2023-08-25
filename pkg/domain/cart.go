package domain




type Cart struct{
	ID uint `json:"id" gorm:"primaryKey"`
	CartID uint `json:"cart_id"`
	Cart Users  `json:"user" gorm:"foreignKey:CartID"`
	ProductID uint `json:"product_id"`
	Product Products `json:"product" gorm:"foreignKey:ProductID"`
	Quantity uint `json:"quantity" `
	Amount float64 `json:"amount"`
}