package repository

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *couponRepository {
	return &couponRepository{
		DB: db,
	}
}

func (repo *couponRepository) AddCoupon(coup models.Coupons) error {
	if err := repo.DB.Exec("INSERT INTO coupons(coupon,discount_rate,valid) values($1,$2,$3)", coup.Coupon, coup.DiscountRate, coup.Valid).Error; err != nil {
		return err
	}

	return nil
}

func (repo *couponRepository) MakeCouponInvalid(id int) error {
	if err := repo.DB.Exec("UPDATE coupons SET valid=false where id=$1", id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *couponRepository) ReActivateCoupon(id int) error {
	if err := repo.DB.Exec("UPDATE coupons SET valid=true where id = $1", id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *couponRepository) FindCouponDetails(couponID int) (domain.Coupons, error) {
	var coupon domain.Coupons
	err := repo.DB.Raw("select * from coupons where id=$1", couponID).Scan(&coupon).Error
	if err != nil {
		return domain.Coupons{}, err
	}

	return domain.Coupons{}, nil
}

func (c *couponRepository) GetAllCoupons() ([]domain.Coupons, error) {
	var model []domain.Coupons
	err := c.DB.Raw("SELECT * FROM coupons").Scan(&model).Error
	if err != nil {
		return []domain.Coupons{}, err
	}

	return model, nil
}
