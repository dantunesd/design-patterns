package handler

import (
	"design-patterns/chain-of-responsiblity/order"
	"design-patterns/chain-of-responsiblity/payment"
	"fmt"
)

type availablePayments map[payment.Payment]uint

type PaymentHandler struct {
	BaseHandler
	available availablePayments
}

func NewPaymentHandler() *PaymentHandler {
	available := availablePayments{
		payment.MONEY:  1,
		payment.CREDIT: 2,
	}

	return &PaymentHandler{
		available: available,
	}
}

func (h *PaymentHandler) Handle(order *order.Order) {
	if !h.isPaymentAvailable(order) {
		h.printUnavailablePayment(order)
		return
	}
	h.BaseHandler.Handle(order)
}

func (h *PaymentHandler) isPaymentAvailable(order *order.Order) bool {
	_, exists := h.available[order.Payment]
	return exists
}

func (h *PaymentHandler) printUnavailablePayment(order *order.Order) {
	fmt.Printf("payment '%s' is unavailable right now.\n", order.Payment)
}
