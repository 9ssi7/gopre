package aggregates

import (
	"errors"
	"time"

	"github.com/9ssi7/gopre/basic-single/internal/domain/entities"
	"github.com/9ssi7/gopre/basic-single/internal/domain/valobj"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order represents an order placed by a customer.
type Order struct {
	ID           uuid.UUID            `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CustomerID   uuid.UUID            `gorm:"not null;type:uuid"`
	Status       valobj.Status        `gorm:"not null"`
	OrderItems   []entities.OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ShippingInfo valobj.ShippingInfo  `gorm:"embedded"`
	BillingInfo  valobj.BillingInfo   `gorm:"embedded"`
	TotalAmount  decimal.Decimal      `gorm:"not null"`
	CreatedAt    time.Time            `gorm:"autoCreateTime"`
	UpdatedAt    time.Time            `gorm:"autoUpdateTime"`
}

// NewOrder creates a new Order aggregate.
func NewOrder(customerID uuid.UUID, items []entities.OrderItem, shippingInfo valobj.ShippingInfo, billingInfo valobj.BillingInfo) (*Order, error) {
	if customerID == uuid.Nil {
		return nil, errors.New("customerID cannot be empty")
	}
	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	totalAmount := decimal.NewFromFloat(0)
	for _, item := range items {
		totalAmount = totalAmount.Add(item.Price.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	return &Order{
		ID:           uuid.New(),
		CustomerID:   customerID,
		Status:       valobj.StatusPending,
		OrderItems:   items,
		ShippingInfo: shippingInfo,
		BillingInfo:  billingInfo,
		TotalAmount:  totalAmount,
	}, nil
}

// AddOrderItem adds a new item to the order.
func (o *Order) AddOrderItem(productID uuid.UUID, quantity int, price decimal.Decimal) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}
	if price.LessThanOrEqual(decimal.NewFromFloat(0)) {
		return errors.New("price must be greater than zero")
	}

	o.OrderItems = append(o.OrderItems, entities.OrderItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	})
	o.TotalAmount = o.TotalAmount.Add(price.Mul(decimal.NewFromInt(int64(quantity))))
	return nil
}

// UpdateShippingInfo updates the shipping information of the order.
func (o *Order) UpdateShippingInfo(shippingInfo valobj.ShippingInfo) error {
	o.ShippingInfo = shippingInfo
	return nil
}

// UpdateBillingInfo updates the billing information of the order.
func (o *Order) UpdateBillingInfo(billingInfo valobj.BillingInfo) error {
	o.BillingInfo = billingInfo
	return nil
}

// Cancel cancels the order.
func (o *Order) Cancel() error {
	if o.Status != valobj.StatusPending && o.Status != valobj.StatusProcessing {
		return errors.New("only pending or processing orders can be cancelled")
	}
	o.Status = valobj.StatusCancelled
	return nil
}

// Ship marks the order as shipped.
func (o *Order) Ship() error {
	if o.Status != valobj.StatusProcessing {
		return errors.New("only processing orders can be shipped")
	}
	o.Status = valobj.StatusShipped
	return nil
}

// Deliver marks the order as delivered.
func (o *Order) Deliver() error {
	if o.Status != valobj.StatusShipped {
		return errors.New("only shipped orders can be delivered")
	}
	o.Status = valobj.StatusDelivered
	return nil
}
