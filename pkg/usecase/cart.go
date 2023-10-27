package usecase

import (
	"errors"
	interfaces "jerseyhub/pkg/repository/interface"
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUseCase         services.UserUseCase
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase) *cartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUseCase:         userUseCase,
	}
}

func (i *cartUseCase) AddToCart(userID, inventoryID int) error {

	//check if item already added if already present send error as already added

	//check if the desired product has quantity available
	stock, err := i.inventoryRepository.CheckStock(inventoryID)
	if err != nil {
		return err
	}
	//if available then call userRepository
	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id
	cart_id, err := i.repo.GetCartId(userID)
	if err != nil {
		return errors.New("some error in geting user cart")
	}
	//if user has no existing cart create new cart
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(userID)
		if err != nil {
			return errors.New("cannot create cart fro user")
		}
	}

	exists, err := i.repo.CheckIfItemIsAlreadyAdded(cart_id, inventoryID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("item already exists in cart")
	}

	//add product to line items
	if err := i.repo.AddLineItems(cart_id, inventoryID); err != nil {
		return errors.New("error in adding products")
	}

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

	products, err := i.userUseCase.GetCart(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	var discountedPrice, totalPrice float64
	for _, v := range products.Data {
		discountedPrice += v.DiscountedPrice
		totalPrice += v.Total
	}

	var checkout models.CheckOut

	checkout.CartID = products.ID
	checkout.Addresses = address
	checkout.Products = products.Data
	checkout.PaymentMethods = payment
	checkout.TotalPrice = totalPrice
	checkout.DiscountedPrice = discountedPrice

	return checkout, err
}
