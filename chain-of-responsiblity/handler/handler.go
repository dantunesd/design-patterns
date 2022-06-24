package handler

import (
	"design-patterns/chain-of-responsiblity/order"
	"fmt"
)

type Handler interface {
	SetNext(next Handler)
	Handle(order *order.Order)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) {
	h.next = next
}

func (h *BaseHandler) Handle(order *order.Order) {
	if h.next != nil {
		h.next.Handle(order)
		return
	}

	h.printEverythingIsOk(order)
}

func (h *BaseHandler) printEverythingIsOk(order *order.Order) {
	fmt.Printf("everything is ok with '%s' order.\n", order.Product)
}
