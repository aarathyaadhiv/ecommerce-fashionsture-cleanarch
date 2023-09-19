package repository

import (
	"time"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type CouponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &CouponRepository{DB}
}

func (c *CouponRepository) IsExist(couponId string) (bool, error) {
	var count int
	err := c.DB.Raw(`SELECT COUNT(*) FROM coupons WHERE coupon_id=? `, couponId).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (c *CouponRepository) IsExpired(coupon string) (bool, error) {
	var expiry time.Time
	err := c.DB.Raw(`SELECT expiry FROM coupons WHERE coupon_id=? `, coupon).Scan(&expiry).Error
	if err != nil {
		return false, err
	}

	return expiry.Unix() < time.Now().Unix(), nil
}

func (c *CouponRepository) CouponDetails(couponId string) (domain.Coupon, error) {
	var coupon domain.Coupon
	err := c.DB.Raw(`SELECT * FROM coupons WHERE coupon_id=?`, couponId).Scan(&coupon).Error
	if err != nil {
		return domain.Coupon{}, err
	}
	return coupon, nil
}

func (c *CouponRepository) AddCoupon(coupon models.AddCoupon) error {
	return c.DB.Exec(`INSERT INTO coupons(coupon_id,discount,usage,expiry,minimum_purchase,maximum_amount) VALUES(?,?,?,?,?,?)`, coupon.CouponId, coupon.Discount, coupon.Usage, coupon.ExpiryTime, coupon.MinimumPurchase, coupon.MaximumAmount).Error
}

func (c *CouponRepository) AddUserCoupon(couponId string, userId, count uint) error {
	return c.DB.Exec(`INSERT INTO user_coupons(coupon_id,users_id,usage_count) VALUES(?,?,?)`, couponId, userId, count).Error
}

func (c *CouponRepository) UpdateUserCount(couponId string, userId uint) error {
	return c.DB.Exec(`UPDATE user_coupons SET usage_count=usage_count+1 WHERE coupon_id=? AND users_id=?`, couponId, userId).Error
}

func (c *CouponRepository) IsUserUsed(couponId string, userId uint) (bool, error) {
	var count int
	err := c.DB.Raw(`SELECT COUNT(*) FROM user_coupons WHERE coupon_id=? AND users_id=?`, couponId, userId).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (c *CouponRepository) UsageCount(couponId string, userId uint) (uint, error) {
	var UsageCount uint
	err := c.DB.Raw(`SELECT usage_count FROM user_coupons WHERE coupon_id=? AND users_id=?`, couponId, userId).Scan(&UsageCount).Error
	if err != nil {
		return 0, err
	}
	return UsageCount, nil
}

func (c *CouponRepository) ExpireCoupon(id uint) error {
	return c.DB.Exec(`UPDATE coupons SET expiry=? WHERE id=?`, time.Now(), id).Error
}

func (c *CouponRepository) BlockCoupon(id uint) error {
	return c.DB.Exec(`UPDATE coupons SET block=true WHERE id=?`, id).Error
}

func (c *CouponRepository) UnBlockCoupon(id uint) error {
	return c.DB.Exec(`UPDATE coupons SET block=false WHERE id=?`, id).Error
}

func (c *CouponRepository) GetCoupon(page, count int) ([]domain.Coupon, error) {
	offset := (page - 1) * count
	var coupon []domain.Coupon
	err := c.DB.Raw(`SELECT * FROM coupons limit ? offset ?`, count, offset).Scan(&coupon).Error
	if err != nil {
		return nil, err
	}
	return coupon, nil
}
