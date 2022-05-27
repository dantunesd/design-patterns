package main

type Aparelho struct {
	Nome string
}

func NewAparelho(nome string) *Aparelho {
	return &Aparelho{
		Nome: nome,
	}
}
