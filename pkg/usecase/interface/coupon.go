package interfaces


type CouponUseCase interface{
	AddCoupon(couponId string,discount int,usage uint)error
	ExpireCoupon(id string)error
	BlockCoupon(id string) error
	UnBlockCoupon(id string) error
}
