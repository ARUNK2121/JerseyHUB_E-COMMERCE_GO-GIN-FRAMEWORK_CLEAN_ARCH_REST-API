package models

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}

type NewPaymentMethod struct {
	PaymentMethod string `json:"payment_method"`
}

type Coupons struct {
	Coupon       string `json:"coupon" gorm:"unique;not null"`
	DiscountRate int    `json:"discount_rate" gorm:"not null"`
	Valid        bool   `json:"valid" gorm:"default:true"`
}
