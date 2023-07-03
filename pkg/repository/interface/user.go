package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	UserBlockStatus(email string) (bool, error)
	AddAddress(id int, address models.AddAddress, result bool) error
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error

	GetCart(id int) ([]models.GetCart, error)
	RemoveFromCart(id int) error
	UpdateQuantityAdd(id int) error
	UpdateQuantityLess(id int) error
	CheckIfFirstAddress(id int) bool
}
