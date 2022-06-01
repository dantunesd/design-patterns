package main

type TransporteGerenciador struct {
	Transporte Transporte
}

func NewTransporteGerenciador() *TransporteGerenciador {
	return &TransporteGerenciador{}
}

func (T *TransporteGerenciador) SetTransporte(transporte Transporte) {
	T.Transporte = transporte
}

func (T *TransporteGerenciador) Transportar(pessoa string) {
	T.Transporte.Transportar(pessoa)
}
