package usecase

import (
	domain "jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) *orderUseCase {
	return &orderUseCase{
		orderRepository: repo,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int) error {

	cart, err := i.orderRepository.GetCart(userid)
	if err != nil {
		return err
	}

	var total float64
	for _, v := range cart {
		total = total + v.Total
	}

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, total)
	if err != nil {
		return err
	}

	if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
		return err
	}

	return nil

}

func (i *orderUseCase) CancelOrder(status string, id int) error {

	err := i.orderRepository.CancelOrder(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("ordered")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("shipped")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("delivered")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Pending = pending
	response.Shipped = shipped
	response.Delivered = delivered

	return response, nil

}
