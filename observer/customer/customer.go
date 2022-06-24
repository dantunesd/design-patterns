package customer

type Status string

const (
	active   Status = "active"
	inactive Status = "inactive"
)

type CustomerEvent string

const (
	CREATED CustomerEvent = "created"
	UPDATED CustomerEvent = "updated"
)

type Customer struct {
	name      string
	birthdate string
	status    Status
}

func NewCustomer(name, birthdate string) *Customer {
	return &Customer{
		name:      name,
		birthdate: birthdate,
		status:    active,
	}
}

func (c *Customer) Activate() {
	c.status = active
}

func (c *Customer) Deactivate() {
	c.status = inactive
}
