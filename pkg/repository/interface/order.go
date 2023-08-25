package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"


type OrderRepository interface{
	AddProductToOrder(orderId,ProductId,quantity,userId uint,amount float64)error
	PlaceOrder(addressId,paymentId,userId uint,amount float64)(uint,error)
	ShowOrderDetails(UserId uint,page,count int)([]models.OrderResponse,error)
}