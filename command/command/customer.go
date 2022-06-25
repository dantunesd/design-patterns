package command

type Status string

type Customer struct {
	Name   string
	Status string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name:   name,
		Status: "active",
	}
}

func (c *Customer) Activate() {
	c.Status = "active"
}

func (c *Customer) Deactivate() {
	c.Status = "inactive"
}
