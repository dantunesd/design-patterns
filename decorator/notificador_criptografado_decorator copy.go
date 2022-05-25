package main

import (
	"encoding/base64"
)

type NotificadorCriptografado struct {
	Notificador INotificador
}

func NewNotificadorCriptografado(notificador INotificador) *NotificadorCriptografado {
	return &NotificadorCriptografado{
		Notificador: notificador,
	}
}

func (N NotificadorCriptografado) Notificar(mensagem string) {
	N.Notificador.Notificar(
		base64.StdEncoding.EncodeToString([]byte(mensagem)),
	)
}
