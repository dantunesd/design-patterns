package main

import "fmt"

type OnibusStrategy struct{}

func NewOnibusStrategy() *OnibusStrategy {
	return &OnibusStrategy{}
}

func (O *OnibusStrategy) Transportar(pessoa string) {
	fmt.Printf("transportando de onibus o(a) %s\n", pessoa)
}
