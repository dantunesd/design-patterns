package ordersystem

import (
	"fmt"
	"time"
)

type Postage struct {
	ID    string
	Order *Order
}

func NewPostage(order *Order) *Postage {
	return &Postage{
		Order: order,
		ID:    time.Now().Format(time.RFC3339),
	}
}

func (p *Postage) Post() {
	fmt.Printf("postageID: %s - posted\n", p.ID)
}
