package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails, ref string) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)

	ChangePassword(id int, old string, password string, repassword string) error
	ForgotPasswordSend(phone string) error
	ForgotPasswordVerifyAndChange(model models.ForgotVerify) error
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error

	GetCart(id int) (models.GetCartResponse, error)
	RemoveFromCart(id int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	GetMyReferenceLink(id int) (string, error)
}
