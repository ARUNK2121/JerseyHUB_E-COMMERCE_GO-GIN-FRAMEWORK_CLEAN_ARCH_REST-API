package usecase

import (
	"errors"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
)

type wishlistUseCase struct {
	repository interfaces.WishlistRepository
	offerRepo  interfaces.OfferRepository
}

func NewWishlistUseCase(repo interfaces.WishlistRepository, offer interfaces.OfferRepository) *wishlistUseCase {
	return &wishlistUseCase{
		repository: repo,
		offerRepo:  offer,
	}
}

func (w *wishlistUseCase) AddToWishlist(user_id, inventory_id int) error {

	if err := w.repository.AddToWishlist(user_id, inventory_id); err != nil {
		return errors.New("could not add to wishlist")
	}

	return nil
}

func (w *wishlistUseCase) RemoveFromWishlist(inventory_id int) error {

	if err := w.repository.RemoveFromWishlist(inventory_id); err != nil {
		return errors.New("could not remove from wishlist")
	}

	return nil
}

func (w *wishlistUseCase) GetWishList(id int) ([]models.Inventories, error) {

	productDetails, err := w.repository.GetWishList(id)
	if err != nil {
		return []models.Inventories{}, err
	}

	//loop inside products and then calculate discounted price of each then return
	for j := range productDetails {
		discount_percentage, err := w.offerRepo.FindDiscountPercentage(productDetails[j].CategoryID)
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
