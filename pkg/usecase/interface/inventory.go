package interfaces

import (
	"jerseyhub/pkg/utils/models"
	"mime/multipart"
)

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventories, image *multipart.FileHeader) (models.InventoryResponse, error)
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	DeleteInventory(id string) error

	ShowIndividualProducts(sku string) (models.Inventories, error)
	ListProducts(page int) ([]models.Inventories, error)
	SearchProducts(key string) ([]models.Inventories, error)

	UpdateProductImage(id int, file *multipart.FileHeader) error
	EditInventoryDetails(int, models.EditInventoryDetails) error
}
