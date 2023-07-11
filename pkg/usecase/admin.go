package usecase

import (
	"errors"

	domain "jerseyhub/pkg/domain"
	helper "jerseyhub/pkg/helper"
	interfaces "jerseyhub/pkg/repository/interface"
	services "jerseyhub/pkg/usecase/interface"
	"jerseyhub/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	access, refresh, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:        adminDetailsResponse,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

}

func (ad *adminUseCase) BlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

// unblock user
func (ad *adminUseCase) UnBlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

func (ad *adminUseCase) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (i *adminUseCase) NewPaymentMethod(inv string) error {

	err := i.adminRepository.NewPaymentMethod(inv)
	if err != nil {
		return err
	}

	return nil

}
