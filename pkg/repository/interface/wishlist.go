package interfaces

import "jerseyhub/pkg/utils/models"

type WishlistRepository interface {
	AddToWishlist(user_id, inventory_id int) error
	RemoveFromWishlist(inventory_id int) error
	GetWishList(id int) ([]models.Inventories, error)
	CheckIfTheItemIsPresentAtWishlist(userID, productID int) (bool, error)
}
