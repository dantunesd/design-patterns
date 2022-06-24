package handler

import (
	"design-patterns/chain-of-responsiblity/order"
	"design-patterns/chain-of-responsiblity/product"
	"fmt"
)

type availableProducts map[product.Product]uint

type ProductHandler struct {
	BaseHandler
	available availableProducts
}

func NewProductHandler() *ProductHandler {
	available := availableProducts{
		product.BIG_MC:    1,
		product.BIG_TASTY: 2,
	}

	return &ProductHandler{
		available: available,
	}
}

func (p *ProductHandler) Handle(order *order.Order) {
	if !p.isProductAvailable(order) {
		p.printUnavailableProduct(order)
		return
	}

	p.BaseHandler.Handle(order)
}

func (p *ProductHandler) isProductAvailable(order *order.Order) bool {
	_, exists := p.available[order.Product]
	return exists
}

func (p *ProductHandler) printUnavailableProduct(order *order.Order) {
	fmt.Printf("product '%s' is unavailable right now.\n", order.Product)
}
