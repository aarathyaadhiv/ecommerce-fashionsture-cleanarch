package usecase

import (
	
	"strconv"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type OrderUseCase struct {
	repo   repository.OrderRepository
	cart   repository.CartRepository
	coupon repository.CouponRepository
}

func NewOrderUseCase(repo repository.OrderRepository, cart repository.CartRepository, coupon repository.CouponRepository) services.OrderUseCase {
	return &OrderUseCase{repo: repo, cart: cart, coupon: coupon}
}

func (c *OrderUseCase) PlaceOrder(addressId, paymentId, userId uint, couponId string) error {
	amount, err := c.cart.TotalAmountInCart(userId)
	if err != nil {
		return err
	}

	ValidCouponExist, err := c.coupon.ExistWithoutExpiry(couponId)
	if err != nil {
		return err
	}

	if ValidCouponExist {
		var couponDetails domain.Coupon

		IsUserUsed, err := c.coupon.IsUserUsed(couponId, userId)
		if err != nil {
			return err
		}
		couponDetails, err = c.coupon.CouponDetails(couponId)
		if err != nil {
			return err
		}
		if IsUserUsed {
			count, err := c.coupon.UsageCount(couponId, userId)
			if err != nil {
				return err
			}
			if count < couponDetails.Usage {
				err := c.coupon.UpdateUserCount(couponId, userId)
				if err != nil {
					return err
				}

				amount = amount - ((amount * float64(couponDetails.Discount)) / 100)

			}
		} else {
			err := c.coupon.AddUserCoupon(couponId, userId, 1)
			if err != nil {
				return err
			}

			amount = amount - ((amount * float64(couponDetails.Discount)) / 100)
			
		}
	}

	id, err := c.repo.PlaceOrder(addressId, paymentId, userId, amount)
	if err != nil {
		return err
	}
	cartProducts, err := c.cart.ProductsInCart(userId)
	if err != nil {
		return err
	}
	for _, ct := range cartProducts {
		err := c.repo.AddProductToOrder(id, ct.ProductId, ct.Quantity, userId, ct.Amount)
		if err != nil {
			return err
		}
	}
	return nil

}

func (c *OrderUseCase) ShowOrderDetails(userId uint, page, count int) ([]models.OrderResponse, error) {
	return c.repo.ShowOrderDetails(userId, page, count)
}

func (c *OrderUseCase) CancelOrder(id string) error {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.repo.CancelOrder(uint(orderId))
}

func (c *OrderUseCase) AdminApproval(id string) error {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.repo.AdminApproval(uint(orderId))
}

func (c *OrderUseCase) ReturnOrder(id string) error {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return c.repo.ReturnOrder(uint(orderId))
}

func (c *OrderUseCase) ShowOrderToAdmin(page, count int) ([]models.OrderDetailsToAdmin, error) {
	return c.repo.OrderDetailsToAdmin(page, count)
}

func (c *OrderUseCase) SearchOrder(id string) (models.OrderDetailsToAdmin, error) {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return models.OrderDetailsToAdmin{}, err
	}
	return c.repo.SearchOrder(uint(orderId))
}
