package main

type PlugComum struct {
	Nome string
}

func NewPlugComum(nome string) *PlugComum {
	return &PlugComum{
		Nome: nome,
	}
}
