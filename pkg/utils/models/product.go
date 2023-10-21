package models

type ProductsReceiver struct {
	CategoryID  string  `json:"Category_id"`
	ProductName string  `json:"product_name"`
	Size_id     uint    `json:"size_id"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type ProductResponse struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	Category    string  `json:"Category"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type SetNewName struct {
	Current string `json:"current"`
	New     string `json:"new"`
}

type MakeOrder struct {
	UserID          int       `json:"user_id"`
	Products        []GetCart `json:"products"`
	AddressID       int       `json:"address_id"`
	PaymentMethodID int       `json:"payment_id"`
	FinalPrice      float64   `json:"final_price"`
}

type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
	PaymentMethodID int `json:"payment_id"`
	CouponID        int `json:"coupon_id"`
}
