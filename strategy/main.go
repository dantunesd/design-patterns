package main

func main() {
	transporteGerenciador := NewTransporteGerenciador()

	transporteGerenciador.SetTransporte(NewAviaoStrategy())
	transporteGerenciador.Transportar("Joaozinho")

	transporteGerenciador.SetTransporte(NewOnibusStrategy())
	transporteGerenciador.Transportar("Joaozinho")

	transporteGerenciador.SetTransporte(NewTaxiStrategy())
	transporteGerenciador.Transportar("Joaozinho")
}
