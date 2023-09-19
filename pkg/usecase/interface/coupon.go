package interfaces

import (
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
)


type CouponUseCase interface{
	AddCoupon(coupon models.AddCoupon)error
	ExpireCoupon(id string)error
	BlockCoupon(id string) error
	UnBlockCoupon(id string) error
	GetCoupon(pages,counts string)([]domain.Coupon,error)
}
