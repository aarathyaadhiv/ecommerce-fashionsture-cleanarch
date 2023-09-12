package models

type ProductUpdate struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type AddProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	Price      float64  `json:"price" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Discount    float64   `json:"discount" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	BrandID     uint   `json:"brand_id" binding:"required"`
}

type AddCategory struct {
	Name string `json:"name" binding:"required"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64   `json:"price"`
	SellingPrice float64 `json:"selling_price" `
	Image       string `json:"image"`
	Discount   float64   `json:"discount"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	Status      string `json:"status"`
}

type AddBrand struct {
	Name string `json:"name" binding:"required"`
}

type Product struct{
	SellingPrice float64 `json:"selling_price"`
	Quantity uint `json:"quantity"`
}

type ProductSearch struct{
	Word string `json:"word" binding:"required"`
}
