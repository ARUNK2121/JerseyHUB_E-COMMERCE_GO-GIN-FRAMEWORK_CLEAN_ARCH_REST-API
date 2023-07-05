package handler

import (
	"fmt"
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	usecase services.PaymentUseCase
}

func NewPaymentHandler(use services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase: use,
	}
}

func (p *PaymentHandler) MakePaymentRazorPay(c *gin.Context) {

	fmt.Println("1")

	orderID := c.Query("id")
	userID := c.Query("user_id")

	fmt.Println("2")

	orderDetail, err := p.usecase.MakePaymentRazorPay(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	fmt.Println("3")
	fmt.Println(orderDetail)

	fmt.Println(orderDetail.Razor_id)
	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	orderID := c.Query("order_id")
	fmt.Println("this is the order id : ", orderID)
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	fmt.Println("paymentID := ", paymentID, " razorID := ", razorID)
	err := p.usecase.VerifyPayment(paymentID, razorID, orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
