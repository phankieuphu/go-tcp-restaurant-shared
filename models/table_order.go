package models

type TableOrders struct {
	TableId   uint    `gorm:"foreignKey:" json:"table_id"`
	DishId    uint    `gorm:"foreignKey" json:"dish_id"`
	Amount    int     `gorm:"not null" json:"amount"`
	DishPrice float64 `gorm:"not null"  json:"dish_price"`
	Status    string  `gorm:"varchar(100)" json:"status"`
}

func (o *TableOrders) TableName() string {
	return "table_orders"
}
