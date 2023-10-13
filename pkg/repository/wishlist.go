package repository

import (
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

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

func (w *wishlistRepository) GetWishList(id int) ([]models.Inventories, error) {

	var productDetails []models.Inventories

	if err := w.DB.Raw("select inventories.id,inventories.category_id,inventories.product_name,inventories.image,inventories.size,inventories.stock,inventories.price from wishlists join inventories on wishlists.inventory_id=inventories.id where user_id=$1", id).Scan(&productDetails).Error; err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil

}
