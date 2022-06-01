package main

import "fmt"

type TaxiStrategy struct{}

func NewTaxiStrategy() *TaxiStrategy {
	return &TaxiStrategy{}
}

func (O *TaxiStrategy) Transportar(pessoa string) {
	fmt.Printf("transportando de taxi o(a) %s\n", pessoa)
}
