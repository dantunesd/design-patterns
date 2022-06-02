package main

import "time"

type Conta struct {
	EstadoAtual    EstadoConta
	DataVencimento time.Time
}

func NewConta(dataVencimento time.Time) *Conta {
	conta := &Conta{}
	conta.EstadoAtual = NewAberto(conta)
	conta.DataVencimento = dataVencimento
	return conta

}

func (C *Conta) Pagar() {
	C.EstadoAtual.Pagar()
}

func (C *Conta) Estado() string {
	return C.EstadoAtual.Estado()
}

func (C *Conta) TrocarEstado(estado EstadoConta) {
	C.EstadoAtual = estado
}

func (C *Conta) Venceu() bool {
	return C.DataVencimento.Before(time.Now())
}
