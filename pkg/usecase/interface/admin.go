package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	NewPaymentMethod(string) error
}
