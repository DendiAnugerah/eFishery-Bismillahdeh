package categoryservice

type CreateCategory struct {
	IdKategori int    `json:"id_kategori"`
	Name       string `json:"name"`
}

type UpdateCategory struct {
	IdKategori int    `json:"id_kategori"`
	Name       string `json:"name"`
}

type DeleteCategory struct {
	IdKategori int `json:"id_kategori"`
}

type FindById struct {
	IdKategori int `json:"id_kategori"`
}
