package product

type Repository interface {
	Create(product *Product) error
	Update(product *Product) error
	Delete(id int) error
	FindAll() ([]*Product, error)
	FindById(id int) (*Product, error)
}
