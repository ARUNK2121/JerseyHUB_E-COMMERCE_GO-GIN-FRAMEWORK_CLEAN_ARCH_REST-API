package usecase

import (
	"errors"
	"fmt"
	domain "jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	services "jerseyhub/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepository  interfaces.OrderRepository
	couponRepository interfaces.CouponRepository
	userUseCase      services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, coup interfaces.CouponRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository:  repo,
		couponRepository: coup,
		userUseCase:      userUseCase,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error {

	cart, err := i.userUseCase.GetCart(userid)
	if err != nil {
		return err
	}
	fmt.Println("cart:", cart)

	var total float64
	for _, v := range cart {
		total = total + v.DiscountedPrice
	}
	fmt.Println("heyy", total)

	//finding discount if any
	DiscountRate := i.couponRepository.FindCouponDiscount(couponID)

	totalDiscount := (total * float64(DiscountRate)) / 100

	total = total - totalDiscount

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, total)
	if err != nil {
		return err
	}

	if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
		return err
	}

	return nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	err := i.orderRepository.CancelOrder(id)
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

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	returned, err := i.orderRepository.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Returned = returned
	response.Delivered = delivered

	return response, nil

}

func (i *orderUseCase) ReturnOrder(id int) error {

	//should check if the order is already returned peoples will misuse this security breach
	// and will get  unlimited money into their wallet
	status, err := i.orderRepository.CheckIfTheOrderIsAlreadyReturned(id)
	if err != nil {
		return err
	}

	if status == "RETURNED" {
		return errors.New("order already returned")
	}

	//should also check if the order is already returned
	//or users will also earn money by returning pending orders by opting COD

	if status != "DELIVERED" {
		return errors.New("user is trying to return an order which is still not delivered")
	}

	//make order as returned order
	if err := i.orderRepository.ReturnOrder(id); err != nil {
		return err
	}

	//find amount to be credited to user
	amount, err := i.orderRepository.FindAmountFromOrderID(id)
	if err != nil {
		return err
	}

	//find the user
	userID, err := i.orderRepository.FindUserIdFromOrderID(id)
	if err != nil {
		return err
	}
	//find if the user having a wallet
	walletID, err := i.orderRepository.FindWalletIdFromUserID(userID)
	if err != nil {
		return err
	}
	//if no wallet create new one
	if walletID == 0 {
		walletID, err = i.orderRepository.CreateNewWallet(userID)
		if err != nil {
			return err
		}
	}
	//credit the amount into users wallet
	if err := i.orderRepository.CreditToUserWallet(amount, walletID); err != nil {
		return err
	}

	return nil

}
