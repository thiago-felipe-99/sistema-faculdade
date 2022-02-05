package logica

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type (
	erro         = *erros.Aplicação
	pessoa       = entidades.Pessoa
	matéria      = entidades.Matéria
	cursoMatéria = entidades.CursoMatéria
	curso        = entidades.Curso
	cpf          = entidades.CPF
	id           = entidades.ID
)

// Lógica representa as operações que se possa fazer com as entidades da
// aplicação.
type Lógica struct {
	Pessoa
	Matéria
	Curso
}

// NovaLógica cria uma Lógica da aplicação.
func NovaLógica(data *data.Data) *Lógica {
	pessoa := Pessoa{Data: data.Pessoa}
	matéria := Matéria{data: data.Matéria}
	curso := Curso{data: data.Curso, matéria: matéria}

	return &Lógica{
		Pessoa:  pessoa,
		Matéria: matéria,
		Curso:   curso,
	}
}
