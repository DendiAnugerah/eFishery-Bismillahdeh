package categoryservice

import "Fish/model/category"

type CategoryService struct {
	repository category.Repository
}

func NewCatService(repository category.Repository) *CategoryService {
	return &CategoryService{repository: repository}
}

func (s *CategoryService) Create(cre *CreateCategory) error {
	NCategory := &category.Category{
		Name: cre.Name,
	}
	return category.Create(s.repository, NCategory)
}

func (s *CategoryService) Update(up *UpdateCategory) error {
	UCategory := &category.Category{
		Id_category: up.Id_category,
		Name:        up.Name,
	}
	return category.Update(s.repository, UCategory)
}

func (s *CategoryService) Delete(du *DeleteCategory) error {
	return category.Delete(s.repository, du.Id_category)
}

func (s *CategoryService) FindAll() (*[]category.Category, error) {
	return category.FindAll(s.repository)
}

func (s *CategoryService) FindById(f *FindById) (*category.Category, error) {
	return category.FindById(s.repository, f.Id_category)
}
