package interfaces

type WishlistRepository interface {
	AddToWishlist(user_id, inventory_id int) error
	RemoveFromWishlist(inventory_id int) error
}
