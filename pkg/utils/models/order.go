package models

import "time"


type OrderDetails struct{
	Id uint `json:"id"`
	User string `json:"user"`
	OrderDate time.Time `json:"order_date"`
	DeliveryDate time.Time `json:"delivery_date"`
	OrderStatus string `json:"Order_status"`
	Total float64 `json:"Total"`
	City string  `json:"city"`
	State string `json:"state"`
	HouseName string `json:"house_name"`
	Pincode string `json:"pincode"`
	PaymentMethod string `json:"payment_method"`
}

type OrderProductDetails struct{
	Product string `json:"product"`
	Description string `json:"description"`
	PricePerProduct float64 `json:"price_per_product"`
	Quantity uint `json:"quantity"`
	ProductPrice float64 `json:"product_price"`
}

type OrderResponse struct{
	
	OrderDetails OrderDetails `json:"order_details"`
	ProductDetails []OrderProductDetails `json:"order_product_details"`
}

type OrderRequest struct{
	AddressId uint `json:"address_id" binding:"required"`
	PaymentId uint `json:"payment_id" binding:"required"`
	CouponId  string `json:"coupon_id" binding:"required"`
}

type CancelOrder struct{
	ProductId uint `json:"product_id"`
	Quantity uint  `json:"quantity"`
}

type OrderDetailsToAdmin struct{
	Id uint `json:"id"`
	User string `json:"user"`
	OrderDate time.Time `json:"order_date"`
	DeliveryDate time.Time `json:"delivery_date"`
	OrderStatus string `json:"Order_status"`
	Total float64 `json:"Total"`
	PaymentStatus string `json:"payment_status"`
	PaymentMethod string `json:"payment_method"`
	Approval bool `json:"approval"`
}