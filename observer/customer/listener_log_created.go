package customer

import (
	"log"
)

type LogCreatedListener struct{}

func NewLogCreatedListener() *LogCreatedListener {
	return &LogCreatedListener{}
}

func (l *LogCreatedListener) Notify(customer *Customer) {
	log.Printf("customer created: %+v\n", customer)
}

func (l *LogCreatedListener) IsSubscribed(event CustomerEvent) bool {
	return event == CREATED
}
