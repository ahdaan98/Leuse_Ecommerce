package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
	IsDeleted    bool   `json:"is_deleted" gorm:"default:false"`
}

type Order struct {
	gorm.Model
	UserID          uint          `json:"user_id" gorm:"not null"`
	Users           User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"price"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string        `json:"payment_status" gorm:"payment_status:4;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID','REFUND IN PROGRESS','RETURNED TO WALLET')"`
}

type OrderResponse struct {
	UserID          uint          `json:"user_id" gorm:"not null"`
	Users           User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	FinalPrice      float64       `json:"price"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string        `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID')"`
}

type OrderItem struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     uint      `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	InventoryID uint      `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

type OrderDetails struct {
	OrderDetails  Order
	PaymentMethod string
}

type OrdersDetails struct {
	ID            int     `json:"id" gorm:"id"`
	Username      string  `json:"name"`
	Address       string  `json:"address"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	Total         float64 `json:"total"`
}

type OrderItemInv struct {
	ID          int     `json:"id" gorm:"id"`
	ProductName string  `json:"product_name"`
	Category_id int     `json:"category_id"`
	Quantity    int     `json:"quantity"`
	Price       int     `json:"price"`
	Total       float64 `json:"total_price"`
	UserID      uint    `json:"user_id"`
	OrderID     uint    `json:"order_id"`
}
