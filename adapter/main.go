package main

import "fmt"

func main() {
	tomada := NewTomadaComum()

	aparelho := NewAparelho("TV smart")
	fmt.Println(
		tomada.EnergizarAparelho(aparelho),
	)

	aparelhoDiferente := NewAparelhoDiferente("TV smart 3 pinos")

	tomadaAdaptador := NewTomadaAdaptador(tomada)
	fmt.Println(
		tomadaAdaptador.EnergizarAparelho(aparelhoDiferente),
	)
}
