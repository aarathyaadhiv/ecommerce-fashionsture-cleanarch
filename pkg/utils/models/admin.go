package models

type AdminSignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
	PhNo     string `json:"phno" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

type AdminDetails struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	PhNo  string `json:"phno"`
}
type AdminTokenResponse struct {
	Token        string
	AdminDetails AdminDetails
}
type AdminSignUpResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	PhNo     string `json:"phno"`
	Password string `json:"password"`
}

type AdminUserResponse struct {
	ID     uint   `json:"id" `
	Name   string `json:"name" `
	Email  string `json:"email"  `
	PhNo   string `json:"mobile_number" `
	Status bool   `json:"status" `
}

type DashboardRevenue struct{
	DayRevenue float64 `json:"day_revenue"`
	MonthRevenue float64 `json:"month_revenue"`
	YearlyRevenue float64  `json:"yearly_revenue"`
}

type DashboardOrders struct{
	CompleteOrder uint `json:"complete_order"`
	PendingOrder uint  `json:"pending_order"`
	CancelledOrder uint  `json:"cancelled_order"`
	TotalOrder uint  `json:"total_order"`
	TotalOrderedUsers uint  `json:"total_ordered_users"`
}

type DashboardAmount struct{
	CreditedAmount float64 `json:"credited_amount"`
	PendingAmount  float64  `json:"pending_amount"`
}

type DashboardUsers struct{
	TotalUsers uint `json:"Total_users"`
	BlockedUsers uint  `json:"Total_block_users"`
	OrderedUsers uint `json:"ordered_users"`
}

type DashboardProduct struct{
	TotalProducts uint `json:"total_products"`
	OutOfStockProducts uint `json:"out_of_stock_products"`
	TopSellingProduct string `json:"top_selling_products"`
}

type Dashboard struct{
	DashboardRevenue DashboardRevenue `json:"dashboard_revenue"`
	DashboardOrders DashboardOrders `json:"dashboard_orders"`
	DashboardAmount DashboardAmount `json:"dashboard_amount"`
	DashboardUsers DashboardUsers `json:"dashboard_users"`
	DashboardProduct DashboardProduct `json:"dashboard_product"`
}

type SalesReport struct{
	TotalSales float64 `json:"total_sales"`
	TotalOrders uint `json:"total_orders"`
	CompleteOrders uint `json:"complete_orders"`
	PendingOrder uint `json:"pending_orders"`
}