package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type OrderRepository interface {
	AddProductToOrder(orderId, ProductId, quantity, userId uint, amount float64) error
	PlaceOrder(addressId, paymentId, userId uint, amount float64, status string) (uint, error)
	ShowOrderDetails(UserId uint, page, count int) ([]models.OrderResponse, error)
	CancelOrder(id uint) error
	AdminApproval(id uint) error
	ReturnOrder(id uint) error
	OrderDetailforPayment(id uint) (string, float64, error)
	UpdatePaymentStatus(status string, orderId uint) error
	OrderDetailsToAdmin(page, count int) ([]models.OrderDetailsToAdmin, error)
	SearchOrder(id uint) (models.OrderDetailsToAdmin, error)
	OrderDetail(id uint) (domain.Order, error)
	AddToWallet(userId uint, amount float64) error
	UpdateWallet(usersId uint, amount float64) error
	IsWalletExist(usersId uint) (bool, error)
	FetchAmountInWallet(userId uint) (float64, error)
	AdminApprovalWithStatus(id uint) error
	PaymentUsingWallet(userId uint, amount float64) error
	FilterOrderByApproval(page, count int,approval bool) ([]models.OrderDetailsToAdmin, error) 
	FilterOrderByPaymentStatus(page, count int,status string) ([]models.OrderDetailsToAdmin, error)
	GetWallet(userId uint)(models.GetWallet,error)
	FilterOrderByApprovalAndPaymentStatus(page, count int,status string,approval bool) ([]models.OrderDetailsToAdmin, error) 
}
