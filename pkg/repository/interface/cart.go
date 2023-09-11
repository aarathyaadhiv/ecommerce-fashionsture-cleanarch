package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"



type CartRepository interface{
	AddToCart(CartID,productId ,quantity uint,amount float64)error
	RemoveFromCart(cartId,productId uint)error
	ShowProductInCart(cartId uint)([]models.CartProducts,error)
	QuantityOfProductInCart(cartId,productId uint)(uint,error)
	AmountOfProductInCart(cartId,productId uint)(float64,error)
	TotalAmountInCart(CartID uint)(float64,error)
	UpdateCart(CartID,ProductId,quantity uint,amount float64)error
	PaymentMethods()([]string,error)
	ProductsInCart(cartId uint)([]models.ProductsInCart,error)
	EmptyCart(cartId uint)error
}