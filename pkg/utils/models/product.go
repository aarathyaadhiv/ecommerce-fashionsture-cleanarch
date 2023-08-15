package models

type ProductUpdate struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type AddProduct struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Discount    int    `json:"discount" binding:"required"`
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
	Price       uint   `json:"price"`
	Image       string `json:"image"`
	Discount    int    `json:"discount"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
}

type AddBrand struct {
	Name string `json:"name" binding:"required"`
}
