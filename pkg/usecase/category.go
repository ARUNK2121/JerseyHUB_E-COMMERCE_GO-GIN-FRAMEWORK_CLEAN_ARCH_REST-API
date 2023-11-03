package usecase

import (
	"errors"
	"fmt"
	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"
)

type categoryUseCase struct {
	repository          interfaces.CategoryRepository
	inventoryRepository interfaces.InventoryRepository
	offerRepository     interfaces.OfferRepository
}

func NewCategoryUseCase(repo interfaces.CategoryRepository, inv interfaces.InventoryRepository, offer interfaces.OfferRepository) services.CategoryUseCase {
	return &categoryUseCase{
		repository:          repo,
		inventoryRepository: inv,
		offerRepository:     offer,
	}
}

func (Cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {

	productResponse, err := Cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil

}

func (Cat *categoryUseCase) UpdateCategory(current string, new string) (domain.Category, error) {

	result, err := Cat.repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}

	if !result {
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}

	newcat, err := Cat.repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}

	return newcat, err
}

func (Cat *categoryUseCase) DeleteCategory(categoryID string) error {

	err := Cat.repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil

}

func (Cat *categoryUseCase) GetCategories() ([]domain.Category, error) {

	categories, err := Cat.repository.GetCategories()
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil

}

func (i *categoryUseCase) GetProductDetailsInACategory(id int) ([]models.Inventories, error) {

	productDetails, err := i.inventoryRepository.ListProductsByCategory(id)
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

	}

	return productDetails, nil

}

func (Cat *categoryUseCase) GetBannersForUsers() ([]models.Banner, error) {
	// Find categories with the highest offer percentage, at least one, maximum 3.
	banners, err := Cat.repository.GetBannersForUsers()
	if err != nil {
		return nil, err
	}

	// Find images of 2 products from each category.
	for i := range banners {
		images, err := Cat.repository.GetImagesOfProductsFromACategory(banners[i].CategoryID)
		if err != nil {
			return nil, err
		}
		banners[i].Images = images
		fmt.Println("loop instance", banners[i])
	}

	fmt.Println("banners", banners)
	return banners, nil
}
