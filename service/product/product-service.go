package productservice

import (
	"Fish/model/category"
	"Fish/model/product"
	"fmt"
)

type ProductService struct {
	repository product.Repository
	category   category.Repository
}

func NewProdService(repositoryR product.Repository, categoryR category.Repository) *ProductService {
	return &ProductService{repository: repositoryR, category: categoryR}
}

func (s *ProductService) Create(cre *CreateProduct) error {
	cat, err := category.FindById(s.category, cre.Id_category)
	if err != nil {
		return err
	}
	if cat == nil {
		return fmt.Errorf("kategori tidak ditemukan")
	}

	NProduct := &product.Product{
		Name:        cre.Name,
		Description: cre.Description,
		Price:       cre.Price,
		Id_category: cre.Id_category,
	}
	return product.Create(s.repository, NProduct)
}

func (s *ProductService) Update(up *UpdateProduct) error {
	UProduct := &product.Product{
		Id_product:  up.Id_product,
		Name:        up.Name,
		Description: up.Description,
		Price:       up.Price,
		Id_category: up.Id_category,
	}

	return product.Update(s.repository, UProduct)
}

func (s *ProductService) FindAll() ([]*product.Product, error) {
	return product.FindAll(s.repository)
}

func (s *ProductService) FindById(f *FindById) (*product.Product, error) {
	return product.FindById(s.repository, f.Id_product)
}

func (s *ProductService) Delete(de *DeleteProduct) error {
	return product.Delete(s.repository, de.Id_product)
}
