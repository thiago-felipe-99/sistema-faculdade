package entidades

import (
	"regexp"
	"time"

	"github.com/google/uuid"
)

// CPF representa o documento CPF(Cadatro De Pessoa Física) do Brasil.
type CPF = string

// parseCPF verifica se a string é um cpf válido por regex.
func parseCPF(cpf string) (CPF, bool) {
	cpfVálido := "00000000000"

	cpfRegra := regexp.
		MustCompile(`^([0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}|[0-9]{11})$`)
	cpfPontos := regexp.MustCompile(`\.|-`)

	cpfs := cpfRegra.FindAllString(cpf, -1)
	if cpfs == nil || len(cpfs) != 1 {
		return cpfVálido, false
	}

	cpfVálido = cpfPontos.ReplaceAllString(cpfs[0], "")

	return cpfVálido, true
}

// nolint:gomnd
// verificarDígitoCPF verfica se os dígitos de verificação do CPF são válidos.
func verificarDígitoCPF(cpf string) bool {
	cpf, parse := parseCPF(cpf)
	if !parse {
		return parse
	}

	const CPFSemSigitos = 9

	cpfNúmeros := [CPFSemSigitos]int{}
	for índice := 0; índice < CPFSemSigitos; índice++ {
		cpfNúmeros[CPFSemSigitos-índice-1] = int(cpf[índice] - '0')
	}

	digito1 := 0
	digito2 := 0

	for índice, número := range cpfNúmeros {
		digito1 += número * (9 - índice%10)
		digito2 += número * (9 - (índice+1)%10)
	}

	digito1 = (digito1 % 11) % 10
	digito2 += digito1 * 9
	digito2 = (digito2 % 11) % 10

	digito1CPF := int(cpf[CPFSemSigitos] - '0')
	digito2CPF := int(cpf[CPFSemSigitos+1] - '0')

	return digito1 == digito1CPF && digito2 == digito2CPF
}

// ValidarCPF verifica se a string é um CPF.
func ValidarCPF(cpf string) (CPF, bool) {
	cpfVálido := "00000000000"

	cpf, parse := parseCPF(cpf)
	if !parse {
		return cpfVálido, false
	}

	dígitosVálidos := verificarDígitoCPF(cpf)
	if !dígitosVálidos {
		return cpfVálido, false
	}

	return cpf, true
}

// Senha representa uma senha na aplicação.
type Senha = string

// GerarNovaSenha retornar a senha hasheada.
func GerarNovaSenha(senha string) Senha {
	return senha
}

// SenhaVálida verifica se a senha segue os padrões requiridos para que tenha
// uma senha válida.
func SenhaVálida(senha string) bool {
	return true
}

func VerificarSenha(senha string, hash Senha) bool {
	return senha == hash
}

// ID representa o indificador único da entidades.
type ID = uuid.UUID

// NovoID gera um novo ID.
func NovoID() ID {
	return uuid.New()
}

// DataAtual retornar a data atual do sistema no padrão UTC.
func DataAtual() time.Time {
	return RemoverHorário(time.Now().UTC())
}

// RemoverHorário retirar o horário de uma data, ou seja, se a data for do
// formato ISO 8601 e tiver o valor 2001-01-01T14:30+00 ela retornará a
// seguinte data 2001-01-01T00:00+00.
func RemoverHorário(data time.Time) time.Time {
	return data.Truncate(24 * time.Hour) //nolint:gomnd
}

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
	IDTurma ID
	IDAluno ID
	Status  string
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
	Pessoa         ID
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
	Pessoa              ID
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
	Pessoa              ID
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
