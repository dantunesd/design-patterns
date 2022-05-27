package main

type EnvioEncomendaFacade struct{}

func NewEnvioEncomendaFacade() EnvioEncomendaFacade {
	return EnvioEncomendaFacade{}
}

func (e EnvioEncomendaFacade) EnviarEncomenda(dados string) {
	encomenda := NewEncomenda(dados)
	encomenda.VerificarDados()

	postagem := NewPostagem(encomenda)
	postagem.Postar()

	transportadora := NewTransportadora()
	transportadora.Transportar(postagem)
}
