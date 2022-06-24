package customer

import (
	"fmt"
)

type SendEmailListener struct{}

func NewSendEmailListener() *SendEmailListener {
	return &SendEmailListener{}
}

func (l *SendEmailListener) Notify(customer *Customer) {
	fmt.Printf("Sending the following email: 'Welcome %s'\n", customer.name)
}

func (l *SendEmailListener) IsSubscribed(event CustomerEvent) bool {
	return event == CREATED
}
