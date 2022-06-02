package main

import "fmt"

type Aberto struct {
	Conta *Conta
}

func NewAberto(conta *Conta) *Aberto {
	return &Aberto{
		Conta: conta,
	}
}

func (A *Aberto) Pagar() {
	if A.Conta.Venceu() {
		fmt.Println("Alterando estado da conta para Vencida devido a data")
		A.Conta.TrocarEstado(NewVencido(A.Conta))
		return
	}

	fmt.Println("Alterando estado da conta para Pago")
	A.Conta.TrocarEstado(NewPago(A.Conta))
}

func (A *Aberto) Estado() string {
	return "ABERTO"
}
