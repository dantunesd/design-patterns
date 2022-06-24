package customer

type Listener interface {
	Notify(customer *Customer)
	IsSubscribed(event CustomerEvent) bool
}
