package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type InventoryRepository interface {
	AddInventory(inventory domain.Inventories) (models.InventoryResponse, error)
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(id string) (domain.Inventories, error)
	ListProducts(page int) ([]domain.Inventories, error)
	CheckStock(inventory_id int) (int, error)
	CheckPrice(inventory_id int) (float64, error)
	SearchProducts(key string) ([]domain.Inventories, error)
}
