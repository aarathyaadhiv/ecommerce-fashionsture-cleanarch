package domain

type Products struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name" `
	Description string   `json:"description" `
	Quantity    uint     `json:"quantity"`
	Price       float64  `json:"price"`
	SellingPrice float64 `json:"selling_price"`
	Image       string   `json:"image"`
	Discount    float64  `json:"discount"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BrandID     uint     `json:"brand_id"`
	Brand       Brand    `json:"brand" gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Brand struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" `
}

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" `
}
