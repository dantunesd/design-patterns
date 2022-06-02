package main

type EstadoConta interface {
	Pagar()
	Estado() string
}
