package main

func main() {
	atorVerdadeiro := NewAtor("Jack Chan")
	ator := NewDuble(atorVerdadeiro)

	ator.Correr()
	ator.Pular("grade")
	ator.Pular("carro")
	ator.Pular("montanha")
}
