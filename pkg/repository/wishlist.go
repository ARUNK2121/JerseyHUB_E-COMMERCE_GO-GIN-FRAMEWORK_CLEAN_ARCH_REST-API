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

func (w *wishlistRepository) AddToWishlist(userID, inventoryID int) error {

	err := w.DB.Exec(`
		INSERT INTO wishlists (user_id,inventory_id)
		VALUES ($1,$2)`, userID, inventoryID).Error
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

	if err := w.DB.Raw("select inventories.id,inventories.category_id,inventories.product_name,inventories.image,inventories.size,inventories.stock,inventories.price from wishlists join inventories on wishlists.inventory_id=inventories.id where user_id=$1 and is_deleted = false", id).Scan(&productDetails).Error; err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil

}

func (w *wishlistRepository) CheckIfTheItemIsPresentAtCart(userID, productID int) (bool, error) {

	var result int64

	if err := w.DB.Raw(`SELECT COUNT (*)
	 FROM line_items 
	 JOIN carts ON carts.id = line_items.cart_id
	 JOIN users ON users.id = carts.user_id
	 WHERE users.id = $1
	 AND 
	 line_items.inventory_id = $2;`, userID, productID).Scan(&result).Error; err != nil {
		return false, err
	}

	return result > 0, nil

}

func (w *wishlistRepository) CheckIfTheItemIsPresentAtWishlist(userID, productID int) (bool, error) {

	var result int64

	if err := w.DB.Raw(`SELECT COUNT (*)
	 FROM wishlists 
	 WHERE user_id = $1
	 AND 
	 inventory_id = $2;`, userID, productID).Scan(&result).Error; err != nil {
		return false, err
	}

	return result > 0, nil

}
