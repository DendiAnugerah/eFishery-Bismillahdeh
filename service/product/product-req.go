package productservice

type CreateProduct struct {
	Name        string
	Description string
	Price       float64
	Id_category int
}

type UpdateProduct struct {
	Id_product  int
	Name        string
	Price       float64
	Description string
	Id_category int
}

type DeleteProduct struct {
	Id_product int
}

type FindById struct {
	Id_product int
}
