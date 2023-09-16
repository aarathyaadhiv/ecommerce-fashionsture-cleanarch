package usecase

import (
	"errors"

	"strconv"

	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
)

type PaymentUseCase struct {
	Repo      repository.PaymentRepository
	OrderRepo repository.OrderRepository
}

func NewPaymentUseCase(repo repository.PaymentRepository, orderRepo repository.OrderRepository) interfaces.PaymentUseCase {
	return &PaymentUseCase{Repo: repo, OrderRepo: orderRepo}
}

func (c *PaymentUseCase) RazorPayPayment(orderId string) (models.DetailsforPayment, error) {

	ordersId, err := strconv.Atoi(orderId)
	if err != nil {
		return models.DetailsforPayment{}, err
	}
	userName, Total, err := c.OrderRepo.OrderDetailforPayment(uint(ordersId))
	if err != nil {
		return models.DetailsforPayment{}, err
	}

	client := razorpay.NewClient("rzp_test_AVJtv4tbQM9Bps", "AnuJPpqQQqzPnPcoeBtFSKRQ")

	data := map[string]interface{}{
		"amount":   int(Total * 100),
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)

	if err != nil {

		return models.DetailsforPayment{}, err
	}
	razorPayId := body["id"].(string)
	err = c.Repo.AddRazorPayDetails(uint(ordersId), razorPayId)
	if err != nil {
		return models.DetailsforPayment{}, err
	}
	detailsForPayment := models.DetailsforPayment{
		UserName:   userName,
		Total:      Total,
		RazorId:    razorPayId,
		OrderId:    uint(ordersId),
		TotalPrice: int(Total * 100)}
	return detailsForPayment, nil
}

func (c *PaymentUseCase) SaveRazorPayPaymentId(orderId string, signature, paymentId string) error {
	ordersId, err := strconv.Atoi(orderId)
	if err != nil {
		return err
	}

	razor, err := c.Repo.FetchRazorId(uint(ordersId))
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"razorpay_order_id":   razor,
		"razorpay_payment_id": paymentId,
	}

	secret := "AnuJPpqQQqzPnPcoeBtFSKRQ"
	isValid := utils.VerifyPaymentSignature(params, signature, secret)
	if isValid {
		err = c.Repo.UpdatePayment(uint(ordersId), razor, paymentId)
		if err != nil {
			return err
		}
		return c.OrderRepo.UpdatePaymentStatus("paid", uint(ordersId))
	}

	return errors.New("payment is not received from an authentic resource")

}
