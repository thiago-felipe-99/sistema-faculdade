package logica

import "thiagofelipe.com.br/sistema-faculdade-backend/data"

type Lógica struct {
	Pessoa
}

func NovaLógica(data data.Data) Lógica {
	return Lógica{
		Pessoa: Pessoa{data: data.Pessoa},
	}
}
