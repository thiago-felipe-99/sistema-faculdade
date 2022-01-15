package logica

import "thiagofelipe.com.br/sistema-faculdade-backend/data"

// Lógica representa as operações que se possa fazer com as entidades da
// aplicação.
type Lógica struct {
	Pessoa
}

// NovaLógica cria uma Lógica da aplicação.
func NovaLógica(data *data.Data) *Lógica {
	return &Lógica{
		Pessoa: Pessoa{data: data.Pessoa},
	}
}
