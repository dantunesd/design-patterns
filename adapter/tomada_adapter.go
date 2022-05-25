package main

import "fmt"

type TomadaAdapter struct {
	TomadaComum ITomadaComum
}

func NewTomadaAdapter(tomadaComum ITomadaComum) *TomadaAdapter {
	return &TomadaAdapter{
		TomadaComum: tomadaComum,
	}
}

func (t TomadaAdapter) EnergizarAparelho(plug *PlugTresPinos) string {
	fmt.Println("adaptando plug de 3 pinos para usar na tomada Comum")

	plugComum := &PlugComum{Nome: plug.Nome}

	fmt.Println("adaptado")

	return t.TomadaComum.EnergizarAparelho(plugComum)
}
