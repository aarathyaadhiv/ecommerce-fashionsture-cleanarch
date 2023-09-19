package models

import "time"

type AddCoupon struct {
	CouponId        string    `json:"coupon_id"`
	Discount        int       `json:"discount"`
	Usage           uint      `json:"usage"`
	Expiry          string    `json:"expiry"`
	ExpiryTime      time.Time `json:"-"`
	MinimumPurchase float64   `json:"minimum_purchase"`
	MaximumAmount   float64   `json:"maximum_amount"`
}
