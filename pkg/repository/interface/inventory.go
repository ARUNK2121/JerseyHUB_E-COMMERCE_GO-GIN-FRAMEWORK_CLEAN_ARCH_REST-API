package interfaces

import (
	"jerseyhub/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory models.AddInventories, url string) (models.InventoryResponse, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (models.Inventories, error)
	ListProducts(page int) ([]models.Inventories, error)
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	SearchProducts(key string) ([]models.Inventories, error)
	UpdateProductImage(int, string) error
}
