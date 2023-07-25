package usecase

import (
	"errors"
	interfaces "jerseyhub/pkg/repository/interface"
)

type wishlistUseCase struct {
	repository interfaces.WishlistRepository
}

func NewWishlistUseCase(repo interfaces.WishlistRepository) *wishlistUseCase {
	return &wishlistUseCase{
		repository: repo,
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
