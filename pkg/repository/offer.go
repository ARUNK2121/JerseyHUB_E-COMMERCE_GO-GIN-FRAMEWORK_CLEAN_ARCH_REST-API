package repository

import (
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) *offerRepository {
	return &offerRepository{
		DB: db,
	}
}

func (repo *offerRepository) AddNewOffer(model models.OfferMaking) error {
	if err := repo.DB.Exec("INSERT INTO offers(category_id,discount_rate) values($1,$2)", model.CategoryID, model.Discount).Error; err != nil {
		return err
	}

	return nil
}

func (repo *offerRepository) MakeOfferExpire(id int) error {
	if err := repo.DB.Exec("UPDATE offers SET valid=false where id=$1", id).Error; err != nil {
		return err
	}

	return nil
}
