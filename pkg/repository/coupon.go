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

func (repo *couponRepository) FindCouponDiscount(couponID int) int {
	var coupon models.Coupons
	err := repo.DB.Raw("select coupon,discount_rate,valid from coupons where id=$1", couponID).Scan(&coupon).Error
	if err != nil {
		return 0
	}

	return coupon.DiscountRate
}

func (c *couponRepository) GetAllCoupons() ([]domain.Coupons, error) {
	var model []domain.Coupons
	err := c.DB.Raw("SELECT * FROM coupons").Scan(&model).Error
	if err != nil {
		return []domain.Coupons{}, err
	}

	return model, nil
}
