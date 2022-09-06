package productrepository

import (
	"Fish/model/product"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProdRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) Create(c *product.Product) error {
	return p.db.Create(c).Error
}

func (p *ProductRepository) FindAll() ([]*product.Product, error) {
	var products []*product.Product

	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindById(id int) (*product.Product, error) {
	p := new(product.Product)

	err := r.db.Where("id = ?", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) Update(p *product.Product) error {
	return r.db.Save(p).Error
}

func (r *ProductRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&product.Product{}).Error
}
