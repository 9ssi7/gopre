package repos

import (
	"errors"

	"github.com/9ssi7/gopre/basic-single/internal/domain/aggregates"
	"github.com/9ssi7/gopre/basic-single/internal/infra/database"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{db: database.GetDB()}
}

func (r *OrderRepo) FindByID(id string) (*aggregates.Order, error) {
	var order aggregates.Order
	result := r.db.Preload("OrderItems").Preload("ShippingInfo").Preload("BillingInfo").First(&order, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &order, nil
}

func (r *OrderRepo) FindAll() ([]aggregates.Order, error) {
	var orders []aggregates.Order
	result := r.db.Preload("OrderItems").Preload("ShippingInfo").Preload("BillingInfo").Find(&orders)
	return orders, result.Error
}

func (r *OrderRepo) Save(order *aggregates.Order) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(order).Error
}

func (r *OrderRepo) DeleteByID(id string) error {
	result := r.db.Delete(&aggregates.Order{}, "id = ?", id)
	return result.Error
}
