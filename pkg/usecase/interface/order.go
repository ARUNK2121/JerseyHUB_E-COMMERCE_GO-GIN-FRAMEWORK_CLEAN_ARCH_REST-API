package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type OrderUseCase interface {
	GetOrders(id int) ([]domain.OrderDetailsWithImages, error)
	OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders() (domain.AdminOrdersResponse, error)
	ReturnOrder(id int) error
	MakePaymentStatusAsPaid(id int) error
	GetIndividualOrderDetails(id int) (models.IndividualOrderDetails, error)
}
