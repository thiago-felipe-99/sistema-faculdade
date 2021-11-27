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
	Insert(matériaToInsert) (Matéria, errors.Application)
	Update(ID, matériaToInsert) (Matéria, errors.Application)
	Get(ID) (Matéria, errors.Application)
	Delete(ID) errors.Application
}
