package repository

import (
	
	"time"

	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) repository.OrderRepository {
	return &OrderRepository{DB}
}

func (c *OrderRepository) AddProductToOrder(orderId, ProductId, quantity, userId uint, amount float64) error {
	err := c.DB.Exec(`INSERT INTO order_products(order_id,product_id,quantity,amount) VALUES(?,?,?,?)`, orderId, ProductId, quantity, amount).Error
	if err != nil {
		return err
	}
	err = c.DB.Exec(`DELETE FROM carts WHERE cart_id=? AND product_id=?`, userId, ProductId).Error
	if err != nil {
		return err
	}
	err = c.DB.Exec(`UPDATE products SET quantity=quantity-? WHERE id=?`, quantity, ProductId).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderRepository) PlaceOrder(addressId, paymentId, userId uint, amount float64) (uint, error) {
	orderDate := time.Now().Truncate(24 * time.Hour)
	deliveryDate := time.Now().AddDate(0, 0, 7)
	var id uint
	err := c.DB.Raw(`INSERT INTO orders(address_id,payment_id,users_id,order_date,delivery_date,amount,status,payment_status) VALUES(?,?,?,?,?,?,?,?) RETURNING id`, addressId, paymentId, userId, orderDate, deliveryDate, amount, "processing","not paid").Scan(&id).Error
	return id, err
}

func (c *OrderRepository) ShowOrderDetails(UserId uint, page, count int) ([]models.OrderResponse, error) {
	var OrderDetails []models.OrderDetails
	offset := (page - 1) * count
	err := c.DB.Raw(`SELECT o.id,u.name as user,o.order_date,o.delivery_date,o.status as order_status,o.amount as total,a.city,a.state,a.house_name,a.pincode,p.method as payment_method FROM orders AS o JOIN addresses AS a on o.address_id=a.id JOIN users AS u ON o.users_id=u.id JOIN payment_methods AS p ON o.payment_id=p.id WHERE o.users_id=? limit ? offset ?`, UserId, count, offset).Scan(&OrderDetails).Error
	if err != nil {
		return nil, err
	}
	var OrderResponse []models.OrderResponse
	for _, o := range OrderDetails {
		var OrderProduct []models.OrderProductDetails
		err := c.DB.Raw(`SELECT p.name as product,p.description,p.selling_price as price_per_product,o.quantity,o.amount as product_price from order_products AS o JOIN products AS p ON o.product_id=p.id WHERE o.order_id=?`, o.Id).Scan(&OrderProduct).Error
		if err != nil {
			return nil, err
		}
		OrderResponse = append(OrderResponse, models.OrderResponse{ProductDetails: OrderProduct, OrderDetails: o})
	}
	return OrderResponse, nil
}

func (c *OrderRepository) CancelOrder(id uint) error {
	err := c.DB.Exec(`UPDATE orders SET status='cancelled' WHERE id=?`, id).Error
	if err != nil {
		return err
	}
	var cancel []models.CancelOrder
	err = c.DB.Raw(`SELECT product_id,quantity FROM order_products WHERE order_id=?`, id).Scan(&cancel).Error
	if err != nil {
		return err
	}
	for _, can := range cancel {
		err := c.DB.Exec(`UPDATE products SET quantity=quantity + ? WHERE id=?`, can.Quantity, can.ProductId).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *OrderRepository) AdminApproval(id uint) error {
	return c.DB.Exec(`UPDATE orders SET approval=true WHERE id=?`, id).Error
}

func (c *OrderRepository) ReturnOrder(id uint)error{
	return c.DB.Exec(`UPDATE orders SET status='returned' WHERE id=?`,id).Error
}

func (c *OrderRepository) OrderDetailforPayment(id uint)(string,float64,error){
	type Order struct{
		Total float64
		UserName string
	}
	 var order Order
	err:= c.DB.Raw(`SELECT o.amount as total,u.name as user_name FROM orders o JOIN users u ON o.users_id=u.id WHERE o.id=?`,id).Scan(&order).Error
	return order.UserName,order.Total,err
}

func (c *OrderRepository) UpdatePaymentStatus(status string,orderId uint)error{
	
	return c.DB.Exec(`UPDATE orders SET payment_status=? WHERE id=?`,status,orderId).Error
}

func (c *OrderRepository) OrderDetailsToAdmin(page,count int)([]models.OrderDetailsToAdmin,error){
	var OrderDetails []models.OrderDetailsToAdmin
	offset := (page - 1) * count
	err := c.DB.Raw(`SELECT o.id,u.name as user,o.order_date,o.delivery_date,o.status as order_status,o.amount as total,o.payment_status,p.method as payment_method,o.approval FROM orders AS o JOIN users AS u ON o.users_id=u.id JOIN payment_methods AS p ON o.payment_id=p.id  limit ? offset ?`,  count, offset).Scan(&OrderDetails).Error
	if err != nil {
		return nil, err
	}
	return OrderDetails,nil
}

func (c *OrderRepository) SearchOrder(id uint)(models.OrderDetailsToAdmin,error){
	var orderDetails models.OrderDetailsToAdmin
	err:= c.DB.Raw(`SELECT o.id,u.name as user,o.order_date,o.delivery_date,o.status as order_status,o.amount as total,o.payment_status,p.method as payment_method,o.approval FROM orders AS o JOIN users AS u ON o.users_id=u.id JOIN payment_methods AS p ON o.payment_id=p.id  WHERE o.id=?`,id).Scan(&orderDetails).Error
	if err!=nil{
		return models.OrderDetailsToAdmin{},err
	}
	return orderDetails,nil
}