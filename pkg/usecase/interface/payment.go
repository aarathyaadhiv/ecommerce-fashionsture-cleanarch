package interfaces

import "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"


type PaymentUseCase interface{
	RazorPayPayment(orderId string) (models.DetailsforPayment,error)
	SaveRazorPayPaymentId(orderId string,razorId,paymentId string) error 
}