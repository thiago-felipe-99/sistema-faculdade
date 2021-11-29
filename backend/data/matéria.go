package data

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/errors"
)

// Matéria representa a entidade Matéria
type Matéria struct {
	ID                  ID
	Nome                string
	CargaHoráriaSemanal time.Duration
	Créditos            float32
	PréRequisitos       []ID
	Tipo                string
}

// MatériaData representa as opereçãoes que se possa fazer com a entidade
// Matéria
type MatériaData interface {
	Inserir(*Matéria) *errors.Aplicação
	Atualizar(ID, *Matéria) *errors.Aplicação
	Pegar(ID) (*Matéria, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}
