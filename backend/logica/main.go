package logica

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type (
	erro    = *erros.Aplicação
	matéria = entidades.Matéria
	pessoa  = entidades.Pessoa
	cpf     = entidades.CPF
	id      = entidades.ID
)

// Lógica representa as operações que se possa fazer com as entidades da
// aplicação.
type Lógica struct {
	Pessoa
	Matéria
}

// NovaLógica cria uma Lógica da aplicação.
func NovaLógica(data *data.Data) *Lógica {
	return &Lógica{
		Pessoa:  Pessoa{data: data.Pessoa},
		Matéria: Matéria{data: data.Matéria},
	}
}
