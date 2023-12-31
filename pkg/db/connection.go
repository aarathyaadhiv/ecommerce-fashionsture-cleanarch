package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
	domain "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Brand{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderProduct{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.RazorPay{})
	db.AutoMigrate(&domain.Coupon{})
	db.AutoMigrate(&domain.UserCoupon{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.Images{})

	return db, dbErr
}
