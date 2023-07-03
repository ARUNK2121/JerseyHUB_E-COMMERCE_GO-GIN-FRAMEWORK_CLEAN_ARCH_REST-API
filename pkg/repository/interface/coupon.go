package interfaces

import (
	"jerseyhub/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(models.Coupons) error
	MakeCouponInvalid(id int) error
	FindCouponDiscount(couponID int) int
}
