package interfaces


type PaymentRepository interface{
	AddRazorPayDetails(orderId uint,razorId string)error
	UpdatePayment(orderId uint,razorId,paymentId string)error
}