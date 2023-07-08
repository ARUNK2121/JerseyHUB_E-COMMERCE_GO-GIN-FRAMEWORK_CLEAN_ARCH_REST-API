package usecase

import (
	"fmt"
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
	fmt.Println("here 1")
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	fmt.Println("new_id", " ")
	orderDetails.OrderID = newid

	fmt.Println("here 2")

	//get userid
	newuserid, err := strconv.Atoi(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.UserID = newuserid
	fmt.Println("here 3")

	//get username
	username, err := p.repository.FindUsername(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	fmt.Println("here 5")

	//get total
	newfinal, err := p.repository.FindPrice(newid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.FinalPrice = newfinal

	fmt.Println("here 6", newfinal)

	client := razorpay.NewClient("rzp_test_pfmFeCViv6CU5K", "TWCh1tyyZZsIxjYSOmmRrLLg")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	fmt.Println(data)

	fmt.Println("here 7")

	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("error:", err)
		return models.OrderPaymentDetails{}, nil
	}
	fmt.Println("here 8")

	fmt.Println(body)
	razorPayOrderID := body["id"].(string)
	fmt.Println("razor_id", razorPayOrderID)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}

func (p *paymentUsecase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	// to check whether the order is already paid
	// err := p.repository.CheckPaymentStatus(razorID, orderID)
	// if err != nil {
	// 	return err
	// }

	err := p.repository.UpdatePaymentDetails(orderID, paymentID, razorID)
	if err != nil {
		return err
	}

	return nil

}

func (p *paymentUsecase) UseWallet(orderID string, userID string) (models.OrderPaymentDetails, error) {
	fmt.Println("here 1")
	var orderDetails models.OrderPaymentDetails
	//get orderid
	newid, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	fmt.Println("new_id", " ")
	orderDetails.OrderID = newid

	fmt.Println("here 2")

	//get userid
	newuserid, err := strconv.Atoi(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.UserID = newuserid
	fmt.Println("here 3")

	//get username
	username, err := p.repository.FindUsername(newuserid)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}

	orderDetails.Username = username

	fmt.Println("here 5")

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

	fmt.Println("here 6", newfinal)

	client := razorpay.NewClient("rzp_test_pfmFeCViv6CU5K", "TWCh1tyyZZsIxjYSOmmRrLLg")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	fmt.Println(data)

	fmt.Println("here 7")

	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println("error:", err)
		return models.OrderPaymentDetails{}, nil
	}
	fmt.Println("here 8")

	fmt.Println(body)
	razorPayOrderID := body["id"].(string)
	fmt.Println("razor_id", razorPayOrderID)

	orderDetails.Razor_id = razorPayOrderID

	return orderDetails, nil
}
