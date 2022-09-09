package cartservice

import "Fish/model/cart"

type CreateCart struct {
	Cart       []*cart.Cart `json:"cart"`
	Id_product int          `json:"id_product"`
}
