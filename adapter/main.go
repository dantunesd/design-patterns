package main

import "fmt"

func main() {
	tomada := NewTomadaComum()

	aparelho := NewPlugComum("TV smart")
	fmt.Println(
		tomada.EnergizarAparelho(aparelho),
	)

	aparelhoTresPinos := NewPlugTresPinos("TV smart 3 pinos")
	tomadaAdapter := NewTomadaAdapter(tomada)
	fmt.Println(
		tomadaAdapter.EnergizarAparelho(aparelhoTresPinos),
	)
}
