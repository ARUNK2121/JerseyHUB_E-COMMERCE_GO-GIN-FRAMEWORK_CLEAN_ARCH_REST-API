package interfaces

import "jerseyhub/pkg/utils/models"

type CartRepository interface {
	GetCart(id int) ([]models.GetCart, error)
	GetAddresses(id int) ([]models.Address, error)
	GetPaymentOptions() ([]models.PaymentMethod, error)
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	AddLineItems(cart_id, inventory_id int) error
}
