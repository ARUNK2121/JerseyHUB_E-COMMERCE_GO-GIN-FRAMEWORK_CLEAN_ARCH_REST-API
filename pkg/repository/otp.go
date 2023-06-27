package repository

import (
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}

func (ot *otpRepository) FindUserByMobileNumber(phone string) bool {

	var count int
	if err := ot.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (ot *otpRepository) UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error) {

	var usersDetails models.UserDetailsResponse
	if err := ot.DB.Raw("select * from users where phone = ?", phone).Scan(&usersDetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}

	return usersDetails, nil

}
