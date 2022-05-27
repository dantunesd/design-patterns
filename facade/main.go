package main

func main() {
	envioEncomenda := NewEnvioEncomendaFacade()
	envioEncomenda.EnviarEncomenda("encomenda exemplo")
}
