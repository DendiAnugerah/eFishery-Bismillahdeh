package categoryservice

type CreateCategory struct {
	Id_category int
	Name        string
}

type UpdateCategory struct {
	Id_category int
	Name        string
}

type DeleteCategory struct {
	Id_category int
}

type FindById struct {
	Id_category int
}
