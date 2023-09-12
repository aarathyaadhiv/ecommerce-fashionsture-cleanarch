package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"


type CouponUseCase interface{
	AddCoupon(couponId string,discount int,usage uint)error
	ExpireCoupon(id string)error
	BlockCoupon(id string) error
	UnBlockCoupon(id string) error
	GetCoupon()([]domain.Coupon,error)
}
