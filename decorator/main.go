package main

func main() {
	notificador := NewNotificador()
	notificador.Notificar("mensagem padrão")

	notificadorDuplicador := NewNotificadorDuplicador(notificador)
	notificadorDuplicador.Notificar("essa mensagem será duplicada")

	notificadorCriptografado := NewNotificadorCriptografado(notificador)
	notificadorCriptografado.Notificar("mensagem criptografada")

	notificadorCriptografadoDuplicador := NewNotificadorCriptografado(notificadorDuplicador)
	notificadorCriptografadoDuplicador.Notificar("essa mensagem será duplicada e criptografada")
}
