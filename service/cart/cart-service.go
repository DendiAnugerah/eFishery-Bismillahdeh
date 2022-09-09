package cartservice

import (
	model "Fish/model/cart"
	"Fish/model/product"
)

type CartService struct {
	repository  model.Repository
	productRepo product.Repository
}

func NewCartService(repository model.Repository, productRepo product.Repository) *CartService {
	return &CartService{repository: repository, productRepo: productRepo}
}

func (s *CartService) Create(cart *CreateCart) error {
	var totalPrice float64
	for _, p := range cart.Cart {
		product, err := product.FindById(s.productRepo, p.Id_product)
		if err != nil {
			return err
		}
		totalPrice += product.Price * float64(p.Quantity)
	}

	newCart := &model.Cart{
		Quantity: model.Cart.Quantity,

		Id_product: cart.Id_product,
	}
}
