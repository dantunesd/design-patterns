package main

import "design-patterns/facade/ordersystem"

type OrderSenderFacade struct{}

func NewOrderSenderFacade() OrderSenderFacade {
	return OrderSenderFacade{}
}

func (o *OrderSenderFacade) Send(data string) {
	order := ordersystem.NewOrder(data)
	order.VerifyData()

	postage := ordersystem.NewPostage(order)
	postage.Post()

	shippingCompany := ordersystem.NewShippingCompany()
	shippingCompany.Transport(postage)
}
