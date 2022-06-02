package main

import "fmt"

type Pago struct {
	Conta *Conta
}

func NewPago(conta *Conta) *Pago {
	return &Pago{
		Conta: conta,
	}
}

func (P *Pago) Pagar() {
	fmt.Println("A conta ja está paga")
}

func (P *Pago) Estado() string {
	return "PAGO"
}
