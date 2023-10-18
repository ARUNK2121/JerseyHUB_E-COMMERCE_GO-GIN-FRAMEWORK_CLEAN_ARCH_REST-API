package repository

import (
	"jerseyhub/pkg/domain"
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
	if err := repo.DB.Exec("DELETE FROM offers WHERE id = $1", id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *offerRepository) FindDiscountPercentage(cat_id int) (int, error) {
	var percentage int
	err := repo.DB.Raw("select discount_rate from offers where category_id=$1 and valid=true", cat_id).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}

	return percentage, nil
}

func (c *offerRepository) GetOffers() ([]domain.Offer, error) {
	var model []domain.Offer
	err := c.DB.Raw("SELECT * FROM offers").Scan(&model).Error
	if err != nil {
		return []domain.Offer{}, err
	}

	return model, nil
}
