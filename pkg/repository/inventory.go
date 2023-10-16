package repository

import (
	"errors"
	"jerseyhub/pkg/utils/models"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) *inventoryRepository {
	return &inventoryRepository{
		DB: DB,
	}
}

func (i *inventoryRepository) AddInventory(inventory models.AddInventories, url string) (models.InventoryResponse, error) {

	query := `
    INSERT INTO inventories (category_id, product_name, size, stock, price, image)
    VALUES (?, ?, ?, ?, ?, ?);
    `
	i.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price, url)

	var inventoryResponse models.InventoryResponse

	return inventoryResponse, nil

}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {

	// Check the database connection
	if i.DB == nil {
		return models.InventoryResponse{}, errors.New("database connection is nil")
	}

	// Update the
	if err := i.DB.Exec("UPDATE inventories SET stock = stock + $1 WHERE id= $2", stock, pid).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	// Retrieve the update
	var newdetails models.InventoryResponse
	var newstock int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&newstock).Error; err != nil {
		return models.InventoryResponse{}, err
	}
	newdetails.ProductID = pid
	newdetails.Stock = newstock

	return newdetails, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

// detailed product details
func (i *inventoryRepository) ShowIndividualProducts(id string) (models.Inventories, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return models.Inventories{}, errors.New("convertion not happened")
	}
	var product models.Inventories
	err := i.DB.Raw(`
	SELECT
		*
		FROM
			inventories
		
		WHERE
			inventories.id = ?
			`, pid).Scan(&product).Error

	if err != nil {
		return models.Inventories{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (ad *inventoryRepository) ListProducts(page int) ([]models.Inventories, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 5
	var productDetails []models.Inventories

	if err := ad.DB.Raw("select id,category_id,product_name,image,size,stock,price from inventories limit $1 offset $2", 5, offset).Scan(&productDetails).Error; err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil

}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var k int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id=$1", pid).Scan(&k).Error; err != nil {
		return 0, err
	}
	return k, nil
}

func (i *inventoryRepository) CheckPrice(pid int) (float64, error) {
	var k float64
	err := i.DB.Raw("SELECT price FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
		return 0, err
	}

	return k, nil
}

func (ad *inventoryRepository) SearchProducts(key string) ([]models.Inventories, error) {
	var productDetails []models.Inventories

	query := `
		SELECT *
		FROM inventories
		WHERE product_name ILIKE '%' || ? || '%'
	`

	if err := ad.DB.Raw(query, key).Scan(&productDetails).Error; err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil
}

func (i *inventoryRepository) UpdateProductImage(int, string) error {

	err := i.DB.Exec("").Error
	if err != nil {
		return err
	}

	return nil
}
