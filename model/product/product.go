package product

type Product struct {
	Id_product  int     `json:"id_product" gorm:"primary_key"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Id_category int     `json:"id_category" gorm:"foreignKey:Id_category"`
}

func Create(r Repository, p *Product) error {
	return r.Create(p)
}

func Update(r Repository, p *Product) error {
	return r.Update(p)
}

func Delete(r Repository, id int) error {
	return r.Delete(id)
}

func FindById(r Repository, id int) (*Product, error) {
	return r.FindById(id)
}

func FindAll(r Repository) ([]*Product, error) {
	return r.FindAll()
}

//type Products []Product

// type Photo struct {
// 	PhotoName string
// }

// type Cart struct {
// 	IdCart   int
// 	Product  Product
// 	Kategori Kategori
// }

// type Kategori struct {
// 	IdKategoricategory
// 	NamaKategori string
// }

// type FProduct struct {
// 	Id      int
// 	NamaFP  string
// 	Product Product
// }

// type LProduct struct {
// 	IdLProduct int
// 	Harga      int
// 	IdCart     Cart
// 	Photo      Photo
// }
