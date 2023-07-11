package handler

import (
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

	orderID := c.Query("id")
	userID := c.Query("user_id")

	orderDetail, err := p.usecase.MakePaymentRazorPay(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not generate order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}

func (p *PaymentHandler) VerifyPayment(c *gin.Context) {

	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	err := p.usecase.VerifyPayment(paymentID, razorID, orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (p *PaymentHandler) MakePaymentFromWallet(c *gin.Context) {

	orderID := c.Query("order_id")
	userID := c.Query("user_id")

	orderDetail, err := p.usecase.UseWallet(orderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not make payment from wallet", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", orderDetail)
}
