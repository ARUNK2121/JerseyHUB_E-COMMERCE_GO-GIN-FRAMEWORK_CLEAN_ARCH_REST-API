package interfaces

import "jerseyhub/pkg/utils/models"

type WishlistUseCase interface {
	AddToWishlist(userID, InventoryID int) error
	RemoveFromWishlist(id int) error
	GetWishList(id int) ([]models.Inventories, error)
}
