package usecase

import (
	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
)

type couponUseCase struct {
	repository interfaces.CouponRepository
}

func NewCouponUseCase(repo interfaces.CouponRepository) *couponUseCase {
	return &couponUseCase{
		repository: repo,
	}
}

func (coup *couponUseCase) AddCoupon(coupon models.Coupons) error {
	if err := coup.repository.AddCoupon(coupon); err != nil {
		return err
	}

	return nil
}

func (coup *couponUseCase) MakeCouponInvalid(id int) error {
	if err := coup.repository.MakeCouponInvalid(id); err != nil {
		return err
	}

	return nil
}

func (coup *couponUseCase) ReActivateCoupon(id int) error {
	if err := coup.repository.ReActivateCoupon(id); err != nil {
		return err
	}

	return nil
}

func (Cat *couponUseCase) GetAllCoupons() ([]domain.Coupons, error) {

	coupons, err := Cat.repository.GetAllCoupons()
	if err != nil {
		return []domain.Coupons{}, err
	}
	return coupons, nil

}
