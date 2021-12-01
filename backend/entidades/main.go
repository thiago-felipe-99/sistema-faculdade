package entidades

import (
	"time"

	"github.com/google/uuid"
)

// CPF representa o documento CPF(Cadatro De Pessoa Física) do Brasil.
type CPF = string

// Senha representa uma senha na aplicação.
type Senha = string

// ID representa o indificador único da entidades.
type ID = uuid.UUID

// CursosOfertado, quando uma turma é ofericido para um certo curso.
type CursosOfertado struct {
	ID
	Vagas   int
	Período string
}

// Nota representa a nota de um aluno.
type Nota struct {
	Aluno  ID
	Nota   float32
	Status string
}

// Horário representa um intervalo de um dia.
type Horário struct {
	Dia            time.Weekday
	HorarioInicial time.Duration
	HorarioFinal   time.Duration
	Turma          ID
	Observacao     string
}

// CursoMatéria representa as matérias que um curso pode ter.
type CursoMatéria struct {
	IDCurso    ID
	IDMatéria  ID
	Período    string
	Tipo       string
	Status     string
	Observação string
}

// TurmaAluno representa as turmas do aluno.
type TurmaAluno struct {
	ID     ID
	Status string
}

// Pessoa representa a entidade Pessoa.
type Pessoa struct {
	ID               ID
	Nome             string
	CPF              CPF
	DataDeNascimento time.Time
	Senha            Senha
}

// Curso representa a entidade Curso.
type Curso struct {
	ID                ID
	Nome              string
	DataDeInício      time.Time
	DataDeDesativação time.Time
	Matérias          []CursoMatéria
}

// Aluno representa a entidade Aluno.
type Aluno struct {
	ID
	Pessoa
	Matrícula      string
	Curso          ID
	DataDeIngresso time.Time
	DataDeSaída    time.Time
	Período        string
	Status         string
	Turmas         []TurmaAluno
}

// Professor representa a entidade Professor.
type Professor struct {
	ID
	Pessoa
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	Turmas              []TurmaAluno
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

// Administrativo representa a entidade Administrativo.
type Administrativo struct {
	ID
	Pessoa
	Matrícula           string
	DataDeIngresso      time.Time
	DataDeSaída         time.Time
	Status              string
	Grau                string
	CargaHoráriaSemanal time.Duration
	HorárioDeAula       Horário
}

// Matéria representa a entidade Matéria.
type Matéria struct {
	ID                  ID
	Nome                string
	CargaHoráriaSemanal time.Duration
	Créditos            float32
	PréRequisitos       []ID
	Tipo                string
}

// Turma representa a entidade Turma.
type Turma struct {
	ID
	Matéria
	Professores        []ID
	Alunos             []ID
	CursosResponsáveis []ID
	CursosOfertados    []CursosOfertado
	HorárioDasAulas    []Horário
	Notas              []Nota
	DataDeInício       time.Time
	DataDeTérmino      time.Time
}
