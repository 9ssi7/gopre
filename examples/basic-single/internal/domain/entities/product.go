package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Product represents a product in the e-commerce store.
type Product struct {
	ID          uuid.UUID       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string          `gorm:"not null"`
	Description string          `gorm:"type:text"`
	Price       decimal.Decimal `gorm:"not null"`
	Stock       int             `gorm:"not null;default:0"` // Available stock
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime"`
}

// NewProduct creates a new Product entity.
func NewProduct(name string, description string, price decimal.Decimal, stock int) (*Product, error) {
	// Perform validation on product data (e.g., check for empty name, negative price, etc.)

	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}, nil
}
