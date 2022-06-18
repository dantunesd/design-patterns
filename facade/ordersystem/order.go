package ordersystem

import "fmt"

type Order struct {
	data string
}

func NewOrder(data string) *Order {
	return &Order{
		data: data,
	}
}

func (o *Order) VerifyData() {
	fmt.Printf("data verified: %s\n", o.data)
}
