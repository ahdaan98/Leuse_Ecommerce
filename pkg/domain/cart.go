package domain

type Cart struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	User   User   `json:"-" gorm:"foreignkey:UserID"`
}

type LineItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CartID      uint      `json:"cart_id" gorm:"not null"`
	Cart        Cart      `json:"-" gorm:"foreignkey:CartID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID;constraint:OnDelete:CASCADE"`
	Quantity    int       `json:"quantity" gorm:"default:1"`
}