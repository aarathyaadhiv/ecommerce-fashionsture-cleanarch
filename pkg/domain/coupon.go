package domain

import "time"

type Coupon struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	CouponId        string    `json:"coupon_id"`
	Discount        int       `json:"discount"`
	Expiry          time.Time `json:"expiry"`
	MinimumPurchase float64   `json:"minimum_purchase"`
	MaximumAmount   float64   `json:"maximum_amount"`
	Usage           uint      `json:"usage"`
	Block           bool      `json:"block" gorm:"default:false"`
}

type UserCoupon struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	CouponId   string `json:"coupon_id"`
	UsersID    uint   `json:"users_id"`
	Users      Users  `json:"users" gorm:"foreignKey:UsersID"`
	UsageCount uint   `json:"usage_count"`
}
