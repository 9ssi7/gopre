package entities

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderItem represents an item within an order.
type OrderItem struct {
	ID        uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OrderID   uuid.UUID       `gorm:"not null;type:uuid"` // Foreign key to Order
	ProductID uuid.UUID       `gorm:"not null;type:uuid"` // Foreign key to Product
	Quantity  int             `gorm:"not null"`
	Price     decimal.Decimal `gorm:"not null"` // Price at the time of ordering
}
