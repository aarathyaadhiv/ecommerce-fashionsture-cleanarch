package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"



type CartUseCase interface{
	AddToCart(cartId,productId uint)error	
	RemoveFromCart(cartId uint,ProductId string)error
	ShowProductInCart(cartId uint)([]models.CartProducts,error)
	TotalAmountInCart(cartId uint)(float64,error)
	EmptyCart(cartId uint)error
}