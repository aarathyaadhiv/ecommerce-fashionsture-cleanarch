package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"


type OrderUseCase interface{
	PlaceOrder(addressId,paymentId,userId uint)error
	ShowOrderDetails(userId uint,page,count int)([]models.OrderResponse,error)
}