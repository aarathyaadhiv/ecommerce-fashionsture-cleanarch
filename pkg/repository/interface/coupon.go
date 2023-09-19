package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)

type CouponRepository interface {
	IsExist(couponId string) (bool, error)
	IsExpired(coupon string) (bool, error)
	CouponDetails(couponId string) (domain.Coupon, error)
	AddCoupon(coupon models.AddCoupon) error
	AddUserCoupon(couponId string, userId, count uint) error
	UpdateUserCount(couponId string, userId uint) error
	IsUserUsed(couponId string, userId uint) (bool, error)
	UsageCount(couponId string, userId uint) (uint, error)
	ExpireCoupon(id uint) error
	BlockCoupon(id uint) error
	UnBlockCoupon(id uint) error
	GetCoupon(page, count int) ([]domain.Coupon, error)
}
