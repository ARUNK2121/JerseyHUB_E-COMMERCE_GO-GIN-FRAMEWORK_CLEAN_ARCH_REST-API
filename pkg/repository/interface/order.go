package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type OrderRepository interface {
	GetOrders(id int) ([]domain.Order, error)
	GetCart(userid int) ([]models.GetCart, error)
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders(status string) ([]domain.OrderDetails, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (domain.Order, error)
}
