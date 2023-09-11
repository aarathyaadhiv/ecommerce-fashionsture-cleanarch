package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"


type CouponRepository interface{
	IsExist(couponId string) (bool, error) 
	IsExpired(coupon string) (bool, error) 
	CouponDetails(couponId string) (domain.Coupon, error)
	AddCoupon(couponId string, discount int, usage uint) error
	AddUserCoupon(couponId string, userId, count uint) error
	UpdateUserCount(couponId string, userId uint) error 
	IsUserUsed(couponId string, userId uint) (bool, error) 
	UsageCount(couponId string, userId uint) (uint, error)
	ExpireCoupon(id uint) error
	BlockCoupon(id uint) error
	UnBlockCoupon(id uint) error
	UpdateExpiry(couponId string)error
	ExistWithoutExpiry(coupon string)(bool,error)
}