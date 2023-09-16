package repository

import (
	
	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"gorm.io/gorm"
)


type PaymentRepository struct{
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB)interfaces.PaymentRepository{
	return &PaymentRepository{DB}
}

func (c *PaymentRepository) AddRazorPayDetails(orderId uint,razorId string)error{
	return c.DB.Exec(`INSERT INTO razor_pays(order_id,razor_id) VALUES(?,?)`,orderId,razorId).Error
}

func (c *PaymentRepository) UpdatePayment(orderId uint,razorId,paymentId string)error{
	
	return c.DB.Exec(`UPDATE razor_pays SET payment_id=? WHERE order_id=? AND razor_id=? `,paymentId,orderId,razorId).Error
}

func (c *PaymentRepository) FetchRazorId(id uint)(string,error){
	var razorId string
	err:=c.DB.Raw(`SELECT razor_id FROM razor_pays WHERE order_id=?`,id).Scan(&razorId).Error
	if err!=nil{
		return "",err
	}
	return razorId,nil
}