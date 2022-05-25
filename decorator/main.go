package main

func main() {
	notificador := NewNotificador()
	notificador.Notificar("mensagem padrão")

	notificadorDuplicador := NewNotificadorDuplicador(notificador)
	notificadorDuplicador.Notificar("essa mensagem será duplicada")

	notificadorCripto := NewNotificadorCriptografado(notificador)
	notificadorCripto.Notificar("mensagem criptografada")
}
