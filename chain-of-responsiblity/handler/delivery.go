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
	if !h.isDeliveryAvailable() {
		h.printUnavailableDelivery()
		return
	}
	h.BaseHandler.Handle(order)
}

func (h *DeliveryHandler) isDeliveryAvailable() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func (h *DeliveryHandler) printUnavailableDelivery() {
	fmt.Println("our delivery is unavailable right now.")
}
