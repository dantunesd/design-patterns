package main

import "fmt"

type Ator struct {
	Nome string
}

func NewAtor(nome string) Ator {
	return Ator{
		Nome: nome,
	}
}

func (a Ator) Correr() {
	fmt.Printf("%s está correndo\n", a.Nome)
}

func (a Ator) Pular(algo string) {
	fmt.Printf("%s está pulando um(a) %s\n", a.Nome, algo)
}
