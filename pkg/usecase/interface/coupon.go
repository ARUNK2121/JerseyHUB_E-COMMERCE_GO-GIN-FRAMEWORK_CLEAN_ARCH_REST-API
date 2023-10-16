package interfaces

import (
	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"
)

type CouponUsecase interface {
	AddCoupon(coupon models.Coupons) error
	MakeCouponInvalid(id int) error
	GetAllCoupons() ([]domain.Coupons, error)
}
