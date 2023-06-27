package usecase

import (
	"errors"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository) *cartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
	}
}

func (i *cartUseCase) AddToCart(user_id, inventory_id int) error {
	//check if the desired product has quantity available
	stock, err := i.inventoryRepository.CheckStock(inventory_id)
	if err != nil {
		return err
	}
	//if available then call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}
	//find price
	price, err := i.inventoryRepository.CheckPrice(inventory_id)
	if err != nil {
		return err
	}
	quantity := 1
	if err := i.repo.AddToCart(user_id, inventory_id, quantity, price); err != nil {
		return err
	}
	//if no error then return nil
	return nil
}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	payment, err := i.repo.GetPaymentOptions()
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.repo.GetCart(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	var price float64
	for _, v := range products {
		price = price + v.Total
	}

	var checkout models.CheckOut

	checkout.Addresses = address
	checkout.Products = products
	checkout.PaymentMethods = payment
	checkout.TotalPrice = price

	return checkout, err
}
