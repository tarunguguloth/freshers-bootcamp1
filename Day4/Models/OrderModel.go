package Models
type Order struct {
	ID        uint   `json:"id"`
	CustomerId   uint   `json:"customer_id" `
	Customer Customer `gorm:"foreignKey:CustomerId:"`
	ProductId  uint   `json:"product_id" `
	Product Product `gorm:"foreignKey:ProductId:"`
	Quantity  uint      `json:"quantity" `
	Status    string    `json:"status"`
}


func (b *Order) TableName() string {
	return "orders"
}