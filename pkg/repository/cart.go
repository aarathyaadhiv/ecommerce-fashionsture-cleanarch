package repository

import (
	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)



type CartRepository struct{
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB)repository.CartRepository{
	return &CartRepository{DB}
}


func (c *CartRepository) AddToCart(CartID,productId ,quantity uint,amount float64)error{
	return c.DB.Exec(`INSERT INTO carts(cart_id,product_id,quantity,amount) VALUES(?,?,?,?)`,CartID,productId,quantity,amount).Error
}

func (c *CartRepository) RemoveFromCart(cartId,productId uint)error{
	return c.DB.Exec(`DELETE FROM carts WHERE cart_id=? and product_id=?`,cartId,productId).Error
}

func (c *CartRepository) ShowProductInCart(cartId uint)([]models.CartProducts,error){
	var cartProducts []models.CartProducts
	err:= c.DB.Raw(`SELECT p.name,p.description,p.image,c.amount,c.quantity FROM carts AS c JOIN products AS p ON c.product_id=p.id WHERE c.cart_id=? `,cartId).Scan(&cartProducts).Error
	return cartProducts,err
}

func (c *CartRepository) QuantityOfProductInCart(cartId,productId uint)(uint,error){
	var quantity uint
	err:=c.DB.Raw(`SELECT quantity FROM carts WHERE cart_id=? AND product_id=?`,cartId,productId).Scan(&quantity).Error
	return quantity,err
}

func (c *CartRepository) AmountOfProductInCart(cartId,productId uint)(float64,error){
	var amount float64
	err:=c.DB.Raw(`SELECT amount FROM carts WHERE cart_id=?,product_id=?`,cartId,productId).Scan(&amount).Error
	return amount,err
}

func (c *CartRepository) TotalAmountInCart(CartID uint)(float64,error){
	var total float64
	err:=c.DB.Raw(`SELECT SUM(amount) FROM carts WHERE cart_id=?`,CartID).Scan(&total).Error
	return total,err
}

func (c *CartRepository) UpdateCart(CartID,ProductId,quantity uint,amount float64)error{
	return c.DB.Exec(`UPDATE carts SET quantity=?,amount=? WHERE cart_id=? AND product_id=?`,quantity,amount,CartID,ProductId).Error
}

func (c *CartRepository) PaymentMethods()([]string,error){
	var methods []string
	err:=c.DB.Raw(`SELECT method FROM payment_methods`).Scan(&methods).Error
	return methods,err
}

func (c *CartRepository) ProductsInCart(cartId uint)([]models.ProductsInCart,error){
	var products []models.ProductsInCart
	err:=c.DB.Raw(`SELECT product_id,quantity,amount FROM carts WHERE cart_id=?`,cartId).Scan(&products).Error
	return products,err
}
