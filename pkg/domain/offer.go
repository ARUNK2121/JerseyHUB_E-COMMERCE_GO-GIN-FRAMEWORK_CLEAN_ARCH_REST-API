package domain

type Offer struct {
	ID           int      `json:"id" gorm:"unique;not null"`
	CategoryID   int      `json:"category_id"`
	Category     Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	DiscountRate int      `json:"discount_rate"`
	Valid        bool     `gorm:"default:True"`
}
