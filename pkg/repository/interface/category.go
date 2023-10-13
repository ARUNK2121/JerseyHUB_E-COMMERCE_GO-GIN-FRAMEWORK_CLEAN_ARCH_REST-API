package interfaces

import "jerseyhub/pkg/domain"

type CategoryRepository interface {
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCategory(currrent string) (bool, error)
	UpdateCategory(current, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategories() ([]domain.Category, error)
}
