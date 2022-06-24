package main

import (
	"design-patterns/chain-of-responsiblity/handler"
	"design-patterns/chain-of-responsiblity/order"
	"design-patterns/chain-of-responsiblity/payment"
	"design-patterns/chain-of-responsiblity/product"
)

func main() {
	firstHandler := handler.NewPaymentHandler()
	secondHandler := handler.NewProductHandler()
	thirdHandler := handler.NewDeliveryHandler()

	firstHandler.SetNext(secondHandler)
	secondHandler.SetNext(thirdHandler)

	orderOk := order.NewOrder(product.BIG_TASTY, payment.CREDIT)
	unavailableProduct := order.NewOrder(product.MC_CHICKEN, payment.CREDIT)
	unavailablePayment := order.NewOrder(product.BIG_TASTY, payment.PIX)

	firstHandler.Handle(orderOk)
	firstHandler.Handle(unavailableProduct)
	firstHandler.Handle(unavailablePayment)
}
