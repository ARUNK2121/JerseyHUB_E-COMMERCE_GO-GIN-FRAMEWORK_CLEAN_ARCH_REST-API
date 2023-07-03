package domain

import "gorm.io/gorm"

type Coupons struct {
	gorm.Model
	Coupon       string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate" gorm:"not null"`
	Valid        bool   `json:"valid" gorm:"default:true"`
}
