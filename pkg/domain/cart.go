package domain

type CartProducts struct {
	ID          uint        `json:"id"  gorm:"primarykey"`
	UserID      uint        `json:"user_id" gorm:"not null"`
	Users       Users       `json:"-" gorm:"foreignkey:UserID"`
	InventoryID uint        `json:"inventory_id"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    float64     `json:"quantity"`
	TotalPrice  float64     `json:"total_price"`
	Deleted     bool        `gorm:"default:false"`
}
