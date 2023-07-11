package usecase

import (
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUsecase struct {
	repository interfaces.PaymentRepository
}

func NewPaymentUseCase(repo interfaces.PaymentRepository) *paymentUsecase {
	return &paymentUsecase{
		repository: repo,
	}
}

func (p *paymentUsecase) MakePaymentRazorPay(orderID string, userID string) (models.OrderPaymentDetails, error) {
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	//get userid
	newuserid, err := strconv.Atoi(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.UserID = newuserid

	//get username
	username, err := p.repository.FindUsername(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	//get total
	newfinal, err := p.repository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.FinalPrice = newfinal

	client := razorpay.NewClient("rzp_test_pfmFeCViv6CU5K", "TWCh1tyyZZsIxjYSOmmRrLLg")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}

	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}

func (p *paymentUsecase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	err := p.repository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	return nil

}

func (p *paymentUsecase) UseWallet(orderID string, userID string) (models.OrderPaymentDetails, error) {
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = newid

	//get userid
	newuserid, err := strconv.Atoi(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.UserID = newuserid

	//get username
	username, err := p.repository.FindUsername(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	//get total
	newfinal, err := p.repository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	//retrieve wallet of the user

	//check if user have enough balance for the payment of newfinal

	//if have sufficient balance then reduce the finalPrice

	//clear the wallet

	//then as usual pay the remaining amount using razorpay and then record payment details in database

	orderDetails.FinalPrice = newfinal

	client := razorpay.NewClient("rzp_test_pfmFeCViv6CU5K", "TWCh1tyyZZsIxjYSOmmRrLLg")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.OrderPaymentDetails{}, nil
	}

	razorPayOrderID := body["id"].(string)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}
