package repository

import "gorm.io/gorm"

type wishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepository {
	return &wishlistRepository{
		DB: db,
	}
}

func (w *wishlistRepository) AddToWishlist(user_id, inventory_id int) error {

	err := w.DB.Exec(`
		INSERT INTO wishlists (user_id,inventory_id)
		VALUES ($1,$2)`, user_id, inventory_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *wishlistRepository) RemoveFromWishlist(id int) error {

	err := w.DB.Exec("UPDATE wishlists SET is_deleted=$1 WHERE inventory_id=$2", true, id).Error
	if err != nil {
		return err
	}

	return nil

}
