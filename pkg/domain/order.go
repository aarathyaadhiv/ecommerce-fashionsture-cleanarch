package domain

import "time"

type Order struct {
	ID            uint          `json:"id"`
	AddressID     uint          `json:"address_id"`
	Address       Address       `json:"address" gorm:"foreignKey:AddressID"`
	PaymentId     uint          `json:"payment_id"`
	PaymentMethod PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentId"`
	OrderDate     time.Time     `json:"order_date"`
	DeliveryDate  time.Time     `json:"delivery_date"`
	Amount        float64       `json:"amount"`
	Status        string        `json:"status" gorm:"default:processing"`
	PaymentStatus string        `json:"payment_status"`
	Approval      bool          `json:"approval" gorm:"default:false"`
	UsersID       uint			`json:"users_id"`
	Users         Users			`json:"users" gorm:"foreignKey:UsersID"`
}

type OrderProduct struct {
	ID        uint     `json:"id"`
	ProductID uint     `json:"product_id"`
	Product   Products `json:"products" gorm:"foreignKey:ProductID"`
	Quantity  uint     `json:"quantity"`
	OrderId   uint     `json:"order_id"`
	Order     Order    `json:"order" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
	Amount    float64  `json:"amount"`
}
