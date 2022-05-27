package main

import "fmt"

type ITomadaComum interface {
	EnergizarAparelho(aparelho *Aparelho) string
}

type TomadaComum struct{}

func NewTomadaComum() *TomadaComum {
	return &TomadaComum{}
}

func (t TomadaComum) EnergizarAparelho(aparelho *Aparelho) string {
	return fmt.Sprintf("energizando o aparelho %s", aparelho.Nome)
}
