package main

type PlugTresPinos struct {
	Nome string
}

func NewPlugTresPinos(nome string) *PlugTresPinos {
	return &PlugTresPinos{
		Nome: nome,
	}
}
