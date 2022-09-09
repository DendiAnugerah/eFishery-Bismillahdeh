package cart

type Repository interface {
	Create(cart *Cart)
}
