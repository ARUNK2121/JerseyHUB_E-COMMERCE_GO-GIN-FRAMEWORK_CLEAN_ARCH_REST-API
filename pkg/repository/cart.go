package repository

import (
	"fmt"
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (i *cartRepository) AddToCart(id, inventory_id, quantity int, price float64) error {
	err := i.DB.Exec(`
		INSERT INTO cart_products (user_id,inventory_id,quantity,total_price)
		VALUES ($1, $2, $3, $4)
		RETURNING id`, id, inventory_id, quantity, price).Error
	if err != nil {
		return err
	}

	return nil
}

func (ad *cartRepository) GetAddresses(id int) ([]models.Address, error) {

	var addresses []models.Address

	if err := ad.DB.Raw("SELECT * FROM addresses WHERE user_id=$1", id).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}

	return addresses, nil

}

func (ad *cartRepository) GetCart(id int) ([]models.GetCart, error) {

	var cart []models.GetCart

	if err := ad.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS Total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventories.id WHERE user_id=$1", id).Scan(&cart).Error; err != nil {
		return []models.GetCart{}, err
	}
	fmt.Println(cart)
	return cart, nil

}

func (ad *cartRepository) GetPaymentOptions() ([]models.PaymentMethod, error) {

	var payment []models.PaymentMethod

	if err := ad.DB.Raw("SELECT * FROM payment_methods").Scan(&payment).Error; err != nil {
		return []models.PaymentMethod{}, err
	}
	fmt.Println(payment)
	return payment, nil

}
