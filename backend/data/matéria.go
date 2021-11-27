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

type matériaToInsert struct {
	Nome                string
	CargaHoráriaSemanal time.Duration
	Créditos            float32
	PréRequisitos       []ID
	Tipo                string
}

// MatériaData representa as opereçãoes que se possa fazer com a entidade
// Matéria
type MatériaData interface {
	Insert(matériaToInsert) (Matéria, errors.ApplicationError)
	Update(ID, matériaToInsert) (Matéria, errors.ApplicationError)
	Get(ID) (Matéria, errors.ApplicationError)
	Delete(ID) errors.ApplicationError
}
