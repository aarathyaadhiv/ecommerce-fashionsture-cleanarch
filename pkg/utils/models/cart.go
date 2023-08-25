package models



type CartProducts struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Amount float64 `json:"amount"`
	Quantity uint  `json:"quantity"`
}


type ProductsInCart struct{
	ProductId uint `json:"product_id"`
	Quantity uint  `json:"quantity"`
	Amount float64  `json:"amount"`
}