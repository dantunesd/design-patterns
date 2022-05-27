package main

import "fmt"

type Duble struct {
	Ator Ator
}

func NewDuble(ator Ator) Duble {
	return Duble{
		Ator: ator,
	}
}

func (d Duble) Correr() {
	d.Ator.Correr()
}

func (d Duble) Pular(algo string) {

	if algo == "montanha" {
		fmt.Printf("O Duble secretamente está Pulando um(a) %s\n", algo)
		return
	}

	d.Ator.Pular(algo)
}
