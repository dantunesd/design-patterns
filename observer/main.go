package main

import "design-patterns/observer/customer"

func main() {
	eventManager := customer.NewEventManager()
	eventManager.AddListener(customer.NewLogCreatedListener())
	eventManager.AddListener(customer.NewLogUpdatedListener())
	eventManager.AddListener(customer.NewSendEmailListener())

	customerJohn := customer.NewCustomer("John", "2000-01-01")
	eventManager.Notify(customer.CREATED, customerJohn)

	customerJohn.Inactivate()
	eventManager.Notify(customer.INACTIVATED, customerJohn)

	customerJohn.Activate()
	eventManager.Notify(customer.ACTIVATED, customerJohn)
}
