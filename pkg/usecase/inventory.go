package usecase

import (
	"errors"
	"fmt"
	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
)

type inventoryUseCase struct {
	repository interfaces.InventoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository) *inventoryUseCase {
	return &inventoryUseCase{
		repository: repo,
	}
}

func (i *inventoryUseCase) AddInventory(inventory domain.Inventories) (models.InventoryResponse, error) {

	InventoryResponse, err := i.repository.AddInventory(inventory)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	return InventoryResponse, nil

}

func (i *inventoryUseCase) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {

	result, err := i.repository.CheckInventory(pid)
	if err != nil {

		return models.InventoryResponse{}, err
	}

	if !result {
		fmt.Println("2")
		return models.InventoryResponse{}, errors.New("there is no inventory as you mentioned")
	}

	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		fmt.Println("3")
		return models.InventoryResponse{}, err
	}

	return newcat, err
}

func (i *inventoryUseCase) DeleteInventory(inventoryID string) error {

	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryUseCase) ShowIndividualProducts(id string) (domain.Inventories, error) {

	product, err := i.repository.ShowIndividualProducts(id)
	if err != nil {
		return domain.Inventories{}, err
	}

	return product, nil

}

func (i *inventoryUseCase) ListProducts(page int) ([]domain.Inventories, error) {

	productDetails, err := i.repository.ListProducts(page)
	if err != nil {
		return []domain.Inventories{}, err
	}

	return productDetails, nil

}
