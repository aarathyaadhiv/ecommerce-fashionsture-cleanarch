package models

type AddCoupon struct {
	CouponId string `json:"coupon_id"`
	Discount int    `json:"discount"`
	Usage    uint   `json:"usage"`
}
