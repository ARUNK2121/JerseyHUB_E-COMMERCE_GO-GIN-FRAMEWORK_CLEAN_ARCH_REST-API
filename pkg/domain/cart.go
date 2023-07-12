package domain

type Cart struct {
	ID     uint  `json:"id" gorm:"primarykey"`
	UserID uint  `json:"user_id" gorm:"not null"`
	Users  Users `json:"-" gorm:"foreignkey:UserID"`
}

type LineItems struct {
	ID          uint        `json:"id" gorm:"primarykey"`
	CartID      uint        `json:"cart_id" gorm:"not null"`
	Cart        Cart        `json:"-" gorm:"foreignkey:CartID"`
	InventoryID uint        `json:"inventory_id" gorm:"not null"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int         `json:"quantity" gorm:"default:1"`
}
