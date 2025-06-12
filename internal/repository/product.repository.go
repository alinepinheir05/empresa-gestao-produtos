package repository

import (
	"github.com/alinepinheir05/empresa-gestao-produtos/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *ProductRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ProductRepository) Delete(id uint) error {
    return r.db.Delete(&domain.Product{}, id).Error
}

func (r *ProductRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}
