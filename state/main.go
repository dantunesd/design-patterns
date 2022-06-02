package main

import (
	"time"
)

func main() {
	dataFutura := time.Now().AddDate(5, 0, 0)
	conta := NewConta(dataFutura)
	conta.Pagar()
	conta.Pagar()

	dataVencida := time.Now().AddDate(-5, 0, 0)
	contaVencida := NewConta(dataVencida)
	contaVencida.Pagar()
	contaVencida.Pagar()
}
