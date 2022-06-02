package main

import "fmt"

type Vencido struct {
	Conta *Conta
}

func NewVencido(conta *Conta) *Vencido {
	return &Vencido{
		Conta: conta,
	}
}

func (v *Vencido) Pagar() {
	fmt.Println("A conta ja está vencida")
}

func (v *Vencido) Estado() string {
	return "VENCIDO"
}
