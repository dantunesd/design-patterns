package transport

import "fmt"

const TAXI transportType = "taxi"

type TaxiStrategy struct {
	transportType transportType
}

func NewTaxiStrategy() *TaxiStrategy {
	return &TaxiStrategy{
		transportType: TAXI,
	}
}

func (t *TaxiStrategy) Transport(person string) {
	fmt.Printf("transporting %s by taxi\n", person)
}

func (t *TaxiStrategy) GetType() transportType {
	return t.transportType
}
