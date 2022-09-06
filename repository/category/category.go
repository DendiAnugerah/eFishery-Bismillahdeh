package categoryrepository

import (
	"Fish/model/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Create(category *category.Category) error {
	return c.db.Create(category).Error
}

func (c *categoryRepository) FindAll() (*[]category.Category, error) {
	var category *[]category.Category
	err := c.db.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) FindById(id int) (*category.Category, error) {
	C := new(category.Category)
	err := c.db.Where("id = ?", id).First(c).Error
	if err != nil {
		return nil, err
	}
	return C, nil
}

func (c *categoryRepository) Update(category *category.Category) error {
	return c.db.Save(category).Error
}

func (c *categoryRepository) Delete(id int) error {
	return c.db.Where("id = ?", id).Delete(&category.Category{}).Error
}
