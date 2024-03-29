package models


type OrderDetails struct {
	OrderID         int     `json:"order_id" gorm:"column:order_id"`
	AddressID       int     `json:"address_id" gorm:"column:address_id"`
	PaymentMethodID int     `json:"payment_method_id" gorm:"column:payment_method_id"`
	Price           float64 `json:"price" gorm:"column:price"`
	OrderStatus     string  `json:"order_status" gorm:"column:order_status"`
	PaymentStatus   string  `json:"payment_status" gorm:"payment_status:4;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID','REFUND IN PROGRESS','RETURNED TO WALLET')"`
}

type OrderDetailsAdmin struct {
	TotalAmount float64 `gorm:"column:total_amount"`
	ProductName string  `gorm:"column:product_name"`
}

type PaymentMethodResponse struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
}

type CombinedOrderDetails struct {
	OrderId       string  `json:"order_id"`
	FinalPrice    float64 `json:"final_price"`
	OrderStatus   string  `json:"order_status" gorm:"column:order_status"`
	PaymentStatus string  `json:"payment_status" gorm:"default:'NOT PAID'"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	HouseName     string  `json:"house_name" validate:"required"`
	Street        string  `json:"street" validate:"required"`
	City          string  `json:"city" validate:"required"`
	State         string  `json:"state" validate:"required"`
	Pin           string  `json:"pin" validate:"required"`
}

type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `josn:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type ItemDetails struct {
	ProductName string  `json:"product_name"`
	FinalPrice  float64 `json:"final_price"`
	Price       float64 `json:"price" `
	Total       float64 `json:"total_price"`
	Quantity    int     `json:"quantity"`
}