package data

import (
	"time"

	"github.com/google/uuid"
)

// ID representa o indificador único da entidades.
type ID = uuid.UUID

// Horário representa um intervalo de um dia.
type Horário struct {
	Dia            time.Weekday
	HorarioInicial time.Duration
	HorarioFinal   time.Duration
	Turma          ID
	Observacao     string
}
