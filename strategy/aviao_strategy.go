package main

import "fmt"

type AviaoStrategy struct{}

func NewAviaoStrategy() *AviaoStrategy {
	return &AviaoStrategy{}
}

func (A *AviaoStrategy) Transportar(pessoa string) {
	fmt.Printf("transportando de aviao o(a) %s\n", pessoa)
}
