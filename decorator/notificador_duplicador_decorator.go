package main

type NotificadorDuplicador struct {
	Notificador INotificador
}

func NewNotificadorDuplicador(notificador INotificador) *NotificadorDuplicador {
	return &NotificadorDuplicador{
		Notificador: notificador,
	}
}

func (N NotificadorDuplicador) Notificar(mensagem string) {
	N.Notificador.Notificar(mensagem + " " + mensagem)
}
