package main

type AparelhoDiferente struct {
	Nome string
}

func NewAparelhoDiferente(nome string) *AparelhoDiferente {
	return &AparelhoDiferente{
		Nome: nome,
	}
}
