package main

import "fmt"

type TomadaAdaptador struct {
	TomadaComum ITomadaComum
}

func NewTomadaAdaptador(tomadaComum ITomadaComum) *TomadaAdaptador {
	return &TomadaAdaptador{
		TomadaComum: tomadaComum,
	}
}

func (t TomadaAdaptador) EnergizarAparelho(aparelhoDiferente *AparelhoDiferente) string {
	fmt.Println("adaptando aparelho com plug diferente para usar na tomada Comum")

	aparelhoComum := &Aparelho{Nome: aparelhoDiferente.Nome}

	fmt.Println("adaptado")

	return t.TomadaComum.EnergizarAparelho(aparelhoComum)
}
