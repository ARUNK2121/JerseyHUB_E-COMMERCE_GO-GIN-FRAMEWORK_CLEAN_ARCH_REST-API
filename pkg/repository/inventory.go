package repository

import (
	"errors"
	"fmt"
	"jerseyhub/pkg/domain"
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

func (i *inventoryRepository) AddInventory(inventory domain.Inventories) (models.InventoryResponse, error) {

	// var id int

	// err := i.DB.Raw("INSERT INTO inventories (category_id,product_name,size,stock,price) VALUES (?, ?, ?, ?, ?) RETURNING id", inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price).Scan(&id).Error
	// if err != nil {
	// 	return models.InventoryResponse{}, err
	// }
	var id uint
	query := `
    INSERT INTO inventories (category_id, product_name, size, stock, price)
    VALUES (?, ?, ?, ?, ?)
    RETURNING id
    `
	i.DB.Raw(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price).Scan(&id)

	var inventoryResponse models.InventoryResponse
	i.DB.Raw(`
	SELECT
		i.product_name,
		i.stock
		FROM
			 inventories i
		WHERE
			inventories.id = ?
			`, id).Scan(&inventoryResponse)

	return inventoryResponse, nil

}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id=?", pid).Scan(&k).Error
	fmt.Println("i:", k)
	if err != nil {
		return false, err
	}

	if k == 0 {
		return false, err
	}

	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {
	fmt.Println("values:", pid, stock)

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
		fmt.Println("debug:1")
		return models.InventoryResponse{}, err
	}
	newdetails.ProductID = pid
	newdetails.Stock = newstock

	fmt.Println(newdetails)

	return newdetails, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}
	fmt.Println("This is the ID:", id)

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

// detailed product details
func (i *inventoryRepository) ShowIndividualProducts(id string) (domain.Inventories, error) {
	pid, error := strconv.Atoi(id)
	if error != nil {
		return domain.Inventories{}, errors.New("convertion not happened")
	}
	var product domain.Inventories
	err := i.DB.Raw(`
	SELECT
		*
		FROM
			inventories
		
		WHERE
			inventories.id = ?
			`, pid).Scan(&product).Error

	if err != nil {
		return domain.Inventories{}, errors.New("error retrieved record")
	}
	return product, nil

}

func (ad *inventoryRepository) ListProducts(page int, count int) ([]domain.Inventories, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productDetails []domain.Inventories

	if err := ad.DB.Raw("select id,category_id,product_name,size,stock,price from inventories limit ? offset ?", count, offset).Scan(&productDetails).Error; err != nil {
		return []domain.Inventories{}, err
	}

	return productDetails, nil

}

func (i *inventoryRepository) CheckStock(pid int) (int, error) {
	var k int
	err := i.DB.Raw("SELECT stock FROM inventories WHERE id=?", pid).Scan(&k).Error
	if err != nil {
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
