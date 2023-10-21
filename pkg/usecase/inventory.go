package usecase

import (
	"errors"
	"fmt"
	helper_interface "jerseyhub/pkg/helper/interface"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
	"mime/multipart"
)

type inventoryUseCase struct {
	repository         interfaces.InventoryRepository
	offerRepository    interfaces.OfferRepository
	helper             helper_interface.Helper
	wishlistRepository interfaces.WishlistRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, offer interfaces.OfferRepository, h helper_interface.Helper, w interfaces.WishlistRepository) *inventoryUseCase {
	return &inventoryUseCase{
		repository:         repo,
		offerRepository:    offer,
		helper:             h,
		wishlistRepository: w,
	}
}

func (i *inventoryUseCase) AddInventory(inventory models.AddInventories, image *multipart.FileHeader) (models.InventoryResponse, error) {

	url, err := i.helper.AddImageToS3(image)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	//send the url and save it in database
	InventoryResponse, err := i.repository.AddInventory(inventory, url)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {

	result, err := i.repository.CheckInventory(pid)
	if err != nil {

		return models.InventoryResponse{}, err
	}

	if !result {
		return models.InventoryResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return newcat, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (models.Inventories, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return models.Inventories{}, err
	}

	DiscountPercentage, err := i.offerRepository.FindDiscountPercentage(product.CategoryID)
	if err != nil {
		return models.Inventories{}, err
	}

	//make discounted price by calculation
	var discount float64
	if DiscountPercentage > 0 {
		discount = (product.Price * float64(DiscountPercentage)) / 100
	}

	product.DiscountedPrice = product.Price - discount

	return product, nil

}

func (i *inventoryUseCase) ListProducts(page, userID int) ([]models.Inventories, error) {

	productDetails, err := i.repository.ListProducts(page)
	if err != nil {
		return []models.Inventories{}, err
	}

	fmt.Println("product details is:", productDetails)

	//loop inside products and then calculate discounted price of each then return
	for j := range productDetails {
		discount_percentage, err := i.offerRepository.FindDiscountPercentage(productDetails[j].CategoryID)
		if err != nil {
			return []models.Inventories{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64

		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}

		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

		productDetails[j].IfPresentAtWishlist, err = i.wishlistRepository.CheckIfTheItemIsPresentAtWishlist(userID, int(productDetails[j].ID))
		if err != nil {
			return []models.Inventories{}, errors.New("error while checking ")
		}

		productDetails[j].IfPresentAtCart, err = i.wishlistRepository.CheckIfTheItemIsPresentAtCart(userID, int(productDetails[j].ID))
		if err != nil {
			return []models.Inventories{}, errors.New("error while checking ")
		}

	}

	return productDetails, nil

}

func (i *inventoryUseCase) SearchProducts(key string) ([]models.Inventories, error) {

	productDetails, err := i.repository.SearchProducts(key)
	if err != nil {
		return []models.Inventories{}, err
	}

	//loop inside products and then calculate discounted price of each then return
	for j := range productDetails {
		discount_percentage, err := i.offerRepository.FindDiscountPercentage(productDetails[j].CategoryID)
		if err != nil {
			return []models.Inventories{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64

		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}

		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

	}

	return productDetails, nil

}

func (i *inventoryUseCase) UpdateProductImage(id int, file *multipart.FileHeader) error {

	url, err := i.helper.AddImageToS3(file)
	if err != nil {
		return err
	}

	//send the url and save it in database
	err = i.repository.UpdateProductImage(id, url)
	if err != nil {
		return err
	}

	return nil

}

func (i *inventoryUseCase) EditInventoryDetails(id int, model models.EditInventoryDetails) error {

	//send the url and save it in database
	err := i.repository.EditInventoryDetails(id, model)
	if err != nil {
		return err
	}

	return nil

}
