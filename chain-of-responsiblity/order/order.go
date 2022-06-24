package order

import (
	"design-patterns/chain-of-responsiblity/payment"
	"design-patterns/chain-of-responsiblity/product"
)

type Order struct {
	Product product.Product
	Payment payment.Payment
}

func NewOrder(name product.Product, payment payment.Payment) *Order {
	return &Order{
		Product: name,
		Payment: payment,
	}
}
