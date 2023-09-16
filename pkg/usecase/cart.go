package usecase

import (
	"errors"
	"strconv"

	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type CartUseCase struct {
	Repo    repository.CartRepository
	product repository.ProductRepository
}

func NewCartUseCase(repo repository.CartRepository, product repository.ProductRepository) services.CartUseCase {
	return &CartUseCase{Repo: repo, product: product}
}

func (c *CartUseCase) AddToCart(cartId, productId uint) error {
	quantity, err := c.Repo.QuantityOfProductInCart(cartId, productId)
	if err != nil {
		return err
	}
	product, err := c.product.FetchProductDetails(productId)
	if err != nil {
		return err
	}
	productQuantity := product.Quantity - quantity
	if productQuantity > 0 {
		if quantity == 0 {
			return c.Repo.AddToCart(cartId, productId, 1, product.SellingPrice)
		}
		amount, err := c.Repo.AmountOfProductInCart(cartId, productId)
		if err != nil {
			return err
		}
		return c.Repo.UpdateCart(cartId, productId, quantity+1, amount+product.SellingPrice)
	}
	return errors.New("product out of stock")

}

func (c *CartUseCase) RemoveFromCart(cartId uint, productId string) error {
	product, err := strconv.Atoi(productId)
	if err != nil {
		return err
	}
	return c.Repo.RemoveFromCart(cartId, uint(product))
}

func (c *CartUseCase) ShowProductInCart(cartId uint,pages,counts string) ([]models.CartProducts, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	Products, err := c.Repo.ShowProductInCart(cartId,page,count)
	if err != nil {
		return nil, err
	}
	updatedCartProduct:=make([]models.CartProducts,0)
	
	for _,product:=range Products{
		quantity,_:=c.product.Quantity(product.Id)
		if quantity==0{
			product.Status="out of stock"
		}else if quantity==1{
			product.Status="only 1 product remains"
		}else{
			product.Status="in stock"
		}
		updatedCartProduct = append(updatedCartProduct, product)
	}
	return updatedCartProduct, nil
}

func (c *CartUseCase) TotalAmountInCart(cartId uint) (float64, error) {
	return c.Repo.TotalAmountInCart(cartId)
}

func (c *CartUseCase) EmptyCart(cartId uint) error {
	return c.Repo.EmptyCart(cartId)
}
