package models


type DetailsforPayment struct{
	UserName string `json:"user_name"`
	RazorId string  `json:"razor_id"`
	Total float64  `json:"total"`
	OrderId uint   `json:"order_id"`
	TotalPrice int  `json:"total_price"`
}
