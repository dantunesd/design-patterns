package command

type DeactivateCustomer struct {
	customer *Customer
}

func NewDeactivateCustomer(request *Customer) *DeactivateCustomer {
	return &DeactivateCustomer{
		customer: request,
	}
}

func (d *DeactivateCustomer) Execute() {
	println("deactivating customer")

	d.customer.Deactivate()
}
