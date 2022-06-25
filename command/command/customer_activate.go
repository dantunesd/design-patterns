package command

type ActivateCustomer struct {
	customer *Customer
}

func NewActivateCustomer(customer *Customer) *ActivateCustomer {
	return &ActivateCustomer{
		customer: customer,
	}
}

func (a *ActivateCustomer) Execute() {
	println("activating customer")

	a.customer.Activate()

}
