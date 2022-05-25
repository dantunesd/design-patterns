package main

import "fmt"

type ITomadaComum interface {
	EnergizarAparelho(plug *PlugComum) string
}

type TomadaComum struct{}

func NewTomadaComum() *TomadaComum {
	return &TomadaComum{}
}

func (t TomadaComum) EnergizarAparelho(aparelhoComum *PlugComum) string {
	return fmt.Sprintf("energizando o aparelho %s", aparelhoComum.Nome)
}
