package category

type Repository interface {
	Create(category *Category) error
	Update(categroy *Category) error
	Delete(id int) error
	FindById(id int) (*Category, error)
	FindAll() (*[]Category, error)
}
