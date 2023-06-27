package interfaces

import "jerseyhub/pkg/utils/models"

type CartRepository interface {
	AddToCart(id, inventory_id, quantity int, price float64) error
	GetCart(id int) ([]models.GetCart, error)
	GetAddresses(id int) ([]models.Address, error)
	GetPaymentOptions() ([]models.PaymentMethod, error)
}
