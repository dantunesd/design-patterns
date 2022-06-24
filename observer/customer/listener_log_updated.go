package customer

import "log"

type LogUpdatedListener struct{}

func NewLogUpdatedListener() *LogUpdatedListener {
	return &LogUpdatedListener{}
}

func (l *LogUpdatedListener) Notify(customer *Customer) {
	log.Printf("customer %s updated to: %s\n", customer.name, customer.status)
}

func (l *LogUpdatedListener) IsSubscribed(event CustomerEvent) bool {
	return event == UPDATED
}
