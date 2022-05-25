package main

import "fmt"

type INotificador interface {
	Notificar(mensagem string)
}

type Notificador struct{}

func NewNotificador() *Notificador {
	return &Notificador{}
}

func (n Notificador) Notificar(mensagem string) {
	fmt.Println(mensagem)
}
