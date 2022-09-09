package cart

type Cart struct {
	Id_cart     int     `json:"id_cart"`
	Quantity    int     `json:"quantity"`
	Total_price float64 `json:"total_price"`
	Id_product  int     `gorm:"foreignKey:Id_product"`
}
