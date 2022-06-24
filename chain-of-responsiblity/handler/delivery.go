package handler

import (
	"design-patterns/chain-of-responsiblity/order"
	"fmt"
	"math/rand"
	"time"
)

type DeliveryHandler struct {
	BaseHandler
}

func NewDeliveryHandler() *DeliveryHandler {
	return &DeliveryHandler{}
}

func (h *DeliveryHandler) Handle(order *order.Order) {
	if h.isAvailable() {
		h.BaseHandler.Handle(order)
		return
	}

	h.printUnavailableDelivery()
}

func (h *DeliveryHandler) isAvailable() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func (h *DeliveryHandler) printUnavailableDelivery() {
	fmt.Println("our delivery is unavailable right now.")
}
