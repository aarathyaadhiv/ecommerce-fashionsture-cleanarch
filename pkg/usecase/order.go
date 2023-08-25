package usecase

import (
	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)


type OrderUseCase struct{
	repo repository.OrderRepository
	cart repository.CartRepository
}

func NewOrderUseCase(repo repository.OrderRepository,cart repository.CartRepository)services.OrderUseCase{
	return &OrderUseCase{repo: repo,cart: cart}
}

func (c *OrderUseCase) PlaceOrder(addressId,paymentId,userId uint)error{
	amount,err:=c.cart.TotalAmountInCart(userId)
	if err!=nil{
		return err
	}
	id,err:=c.repo.PlaceOrder(addressId,paymentId,userId,amount)
	if err!=nil{
		return err
	}
	cartProducts,err:=c.cart.ProductsInCart(userId)
	if err!=nil{
		return err
	}
	for _,ct:=range cartProducts{
		err:=c.repo.AddProductToOrder(id,ct.ProductId,ct.Quantity,userId,ct.Amount)
		if err!=nil{
			return err
		}
	}
	return nil

}

func (c *OrderUseCase) ShowOrderDetails(userId uint,page,count int)([]models.OrderResponse,error){
return c.repo.ShowOrderDetails(userId,page,count)
}