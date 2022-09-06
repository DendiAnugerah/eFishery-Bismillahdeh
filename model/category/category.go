package category

type Category struct {
	IdKategori int
	Name       string
}

func Create(r Repository, category *Category) error {
	err := r.Create(category)
	return err
}

func Update(r Repository, category *Category) error {
	err := r.Update(category)
	return err
}

func Delete(r Repository, id int) error {
	err := r.Delete(id)
	return err
}

func FindAll(r Repository) (*[]Category, error) {
	category, err := r.FindAll()
	return category, err
}

func FindById(r Repository, id int) (*Category, error) {
	category, err := r.FindById(id)
	return category, err
}
