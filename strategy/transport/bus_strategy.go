package transport

import "fmt"

const BUS transportType = "bus"

type BusStrategy struct {
	transportType transportType
}

func NewBusStrategy() *BusStrategy {
	return &BusStrategy{
		transportType: BUS,
	}
}

func (b *BusStrategy) Transport(person string) {
	fmt.Printf("transporting %s by bus\n", person)
}

func (b *BusStrategy) GetType() transportType {
	return b.transportType
}
