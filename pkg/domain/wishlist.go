package domain

type Wishlist struct {
	ID          uint        `json:"id" gorm:"primarykey"`
	UserID      uint        `json:"user_id" gorm:"not null"`
	Users       Users       `json:"-" gorm:"foreignkey:UserID"`
	InventoryID uint        `json:"inventory_id" gorm:"not null"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	IsDeleted   bool        `json:"is_deleted" gorm:"default:false"`
}
