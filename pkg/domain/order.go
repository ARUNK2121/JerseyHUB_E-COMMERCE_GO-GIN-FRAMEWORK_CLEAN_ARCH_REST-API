package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}

type Order struct {
	gorm.Model
	// OrderId         string        `json:"order_id" gorm:"primaryKey;autoIncrement"`
	UserID          uint          `json:"user_id" gorm:"not null"`
	Users           Users         `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"price"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'ordered';check:order_status IN ('ordered', 'shipped', 'delivered','canceled')"`
}

type OrderItem struct {
	ID          uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint        `json:"order_id"`
	Order       Order       `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	InventoryID uint        `json:"inventory_id"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int         `json:"quantity"`
	TotalPrice  float64     `json:"total_price"`
}

type AdminOrdersResponse struct {
	Pending   []Order
	Shipped   []Order
	Delivered []Order
	Canceled  []Order
}
