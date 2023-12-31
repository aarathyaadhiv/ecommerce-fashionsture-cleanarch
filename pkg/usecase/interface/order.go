package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"


type OrderUseCase interface{
	PlaceOrder(addressId,paymentId,userId uint,couponId string)error
	ShowOrderDetails(userId uint,page,count int)([]models.OrderResponse,error)
	CancelOrder(id string)error
	AdminApproval(id string)error
	ReturnOrder(id string)error
	ShowOrderToAdmin(page,count int,approval,paymentStatus string)([]models.OrderDetailsToAdmin,error)
	SearchOrder(id string)(models.OrderDetailsToAdmin,error)
	GetWallet(userId uint)(models.GetWallet,error)
}