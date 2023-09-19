package usecase

import (
	"errors"

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
	cartProducts, err := c.cart.ProductsInCart(userId)
	if err != nil {
		return err
	}
	if cartProducts == nil {
		return errors.New("no products in cart")
	}
	isCouponExist, err := c.coupon.IsExist(couponId)
	if err != nil {
		return err
	}

	if isCouponExist {
		isnotValid, err := c.coupon.IsExpired(couponId)
		if err != nil {
			return err
		}

		if !isnotValid {
			var couponDetails domain.Coupon
			couponDetails, err = c.coupon.CouponDetails(couponId)
			if err != nil {
				return err
			}

			if amount >= couponDetails.MinimumPurchase {
				IsUserUsed, err := c.coupon.IsUserUsed(couponId, userId)
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
						discount := (amount * float64(couponDetails.Discount)) / 100

						if discount > couponDetails.MaximumAmount {
							amount = amount - couponDetails.MaximumAmount
						} else {
							amount = amount - discount

						}

					}
				} else {
					err := c.coupon.AddUserCoupon(couponId, userId, 1)
					if err != nil {
						return err
					}

					discount := (amount * float64(couponDetails.Discount)) / 100

					if discount > couponDetails.MaximumAmount {
						amount = amount - couponDetails.MaximumAmount
					} else {
						amount = amount - discount

					}

				}
			}
		}

	}

	var status string
	if paymentId == 3 {
		err := c.PaymentUsingWallet(userId, amount)
		if err != nil {
			return err
		}
		status = "paid"
	} else {
		status = "not paid"
	}

	id, err := c.repo.PlaceOrder(addressId, paymentId, userId, amount, status)
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
	orderDetails, err := c.repo.OrderDetail(uint(orderId))

	if err != nil {
		return err
	}
	if orderDetails.Status == "delivered" {
		return errors.New("cannot cancel the order as it is delivered")
	}

	return c.repo.CancelOrder(uint(orderId))
}

func (c *OrderUseCase) AdminApproval(id string) error {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	orderDetails, err := c.repo.OrderDetail(uint(orderId))

	if err != nil {
		return err
	}

	status := orderDetails.Status
	payment := orderDetails.PaymentStatus
	if (status == "returned" && payment == "paid") || (status == "cancelled" && payment == "paid") {
		err := c.AddToWallet(orderDetails.UsersID, orderDetails.Amount)
		if err != nil {
			return err
		}
		return c.repo.AdminApprovalWithStatus(uint(orderId))
	}
	return c.repo.AdminApproval(uint(orderId))
}

func (c *OrderUseCase) ReturnOrder(id string) error {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	orderDetails, err := c.repo.OrderDetail(uint(orderId))

	if err != nil {
		return err
	}
	if orderDetails.Status != "delivered" {
		return errors.New(" The order is still in transit, unable to return order")
	}
	return c.repo.ReturnOrder(uint(orderId))
}

func (c *OrderUseCase) ShowOrderToAdmin(page, count int, approval, paymentStatus string) ([]models.OrderDetailsToAdmin, error) {
	if approval != "" && paymentStatus != "" {
		return c.FilterOrderByApprovalAndPaymentStatus(page, count, approval, paymentStatus)
	}
	if approval != "" {
		var approvalStatus bool
		if approval == "approved" {
			approvalStatus = true
		} else if approval == "not approved" {
			approvalStatus = false
		}
		return c.repo.FilterOrderByApproval(page, count, approvalStatus)
	}
	if paymentStatus != "" {
		return c.repo.FilterOrderByPaymentStatus(page, count, paymentStatus)
	}
	return c.repo.OrderDetailsToAdmin(page, count)
}

func (c *OrderUseCase) SearchOrder(id string) (models.OrderDetailsToAdmin, error) {
	orderId, err := strconv.Atoi(id)
	if err != nil {
		return models.OrderDetailsToAdmin{}, err
	}
	return c.repo.SearchOrder(uint(orderId))
}

func (c *OrderUseCase) AddToWallet(userId uint, amount float64) error {
	isWalletExist, err := c.repo.IsWalletExist(userId)
	if err != nil {
		return err
	}
	if isWalletExist {
		return c.repo.UpdateWallet(userId, amount)
	}
	return c.repo.AddToWallet(userId, amount)
}

func (c *OrderUseCase) PaymentUsingWallet(userId uint, amount float64) error {
	isWalletExist, err := c.repo.IsWalletExist(userId)
	if err != nil {
		return err
	}
	if isWalletExist {
		walletAmount, err := c.repo.FetchAmountInWallet(userId)
		if err != nil {
			return err
		}
		if walletAmount >= amount {
			return c.repo.PaymentUsingWallet(userId, amount)
		}
		return errors.New("wallet has no sufficient balance")
	}
	return errors.New("no wallet for user")
}

func (c *OrderUseCase) FilterOrderByApprovalAndPaymentStatus(page, count int, approval, paymentStatus string) ([]models.OrderDetailsToAdmin, error) {

	var approvalStatus bool
	if approval == "approved" {
		approvalStatus = true
	} else if approval == "not approved" {
		approvalStatus = false
	}
	return c.repo.FilterOrderByApprovalAndPaymentStatus(page, count, paymentStatus, approvalStatus)
}

func (c *OrderUseCase) GetWallet(userId uint) (models.GetWallet, error) {
	return c.repo.GetWallet(userId)
}
