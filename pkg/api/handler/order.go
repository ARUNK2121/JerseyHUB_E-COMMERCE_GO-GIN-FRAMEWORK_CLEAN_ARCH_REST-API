package handler

import (
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"
	"jerseyhub/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

// @Summary		Get Orders
// @Description	user can view the details of the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [get]
func (i *OrderHandler) GetOrders(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetOrders(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Order Now
// @Description	user can order the items that currently in cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			order body	models.Order  true	"id"
// @Param			coupon-id query	int  true	"coupon-id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/check-out/order [post]
func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	idstring := c.Query("coupon-id")
	couponId, err := strconv.Atoi(idstring)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coupon id trouble", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID, couponId); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Order Cancel
// @Description	user can cancel the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id  query  string  true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [delete]
func (i *OrderHandler) CancelOrder(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coonversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.CancelOrder(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully canceled the order", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Update Order Status
// @Description	Admin can change the status of the order
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id  query  string  true	"id"
// @Param			status  query  string  true	"status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders/edit/status [put]
func (i *OrderHandler) EditOrderStatus(c *gin.Context) {

	status := c.Query("status")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coonversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.EditOrderStatus(status, id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully edited the order status", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Admin Orders
// @Description	Admin can view the orders according to status
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders [get]
func (i *OrderHandler) AdminOrders(c *gin.Context) {

	orders, err := i.orderUseCase.AdminOrders()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}
