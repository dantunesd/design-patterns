package main

import "fmt"

type Encomenda struct {
	dados string
}

func NewEncomenda(dados string) Encomenda {
	return Encomenda{dados: dados}
}

func (e Encomenda) VerificarDados() {
	fmt.Println("dados verificados")
}

type Postagem struct {
	ID        string
	Encomenda Encomenda
}

func NewPostagem(Encomenda Encomenda) Postagem {
	return Postagem{
		Encomenda: Encomenda,
		ID:        "example",
	}
}

func (p *Postagem) Postar() {
	fmt.Println("encomenda postada")
}

type Transportadora struct{}

func NewTransportadora() Transportadora {
	return Transportadora{}
}

func (t Transportadora) Transportar(postagem Postagem) {
	fmt.Printf("encomenda %s transportada\n", postagem.ID)
}
