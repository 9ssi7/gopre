package repos

import (
	"errors"

	"github.com/9ssi7/gopre/basic-single/internal/domain/entities"
	"github.com/9ssi7/gopre/basic-single/internal/infra/database"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{db: database.GetDB()}
}

func (r *ProductRepo) FindByID(id string) (*entities.Product, error) {
	var product entities.Product
	result := r.db.First(&product, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepo) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepo) Save(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepo) DeleteByID(id string) error {
	result := r.db.Delete(&entities.Product{}, "id = ?", id)
	return result.Error
}
