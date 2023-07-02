package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory domain.Inventories) (models.InventoryResponse, error)
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	ShowIndividualProducts(sku string) (domain.Inventories, error)
	ListProducts(page int) ([]domain.Inventories, error)
}
