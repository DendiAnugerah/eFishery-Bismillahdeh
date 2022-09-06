package productservice

type CreateProduct struct {
	Name       string `json:"name"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Photo      string `json:"photo"`
	IdKategori int    `json:"id_kategori"`
}

type UpdateProduct struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Deskripsi  string `json:"deskripsi"`
	Harga      int    `json:"harga"`
	Photo      string `json:"photo"`
	IdKategori int    `json:"id_kategori"`
}

type DeleteProduct struct {
	Id int `json:"id"`
}

type FindById struct {
	Id int `json:"id"`
}
