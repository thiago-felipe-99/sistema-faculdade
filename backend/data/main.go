package data

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

type ID = entidades.ID

// Pessoa representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type Pessoa interface {
	Inserir(*entidades.Pessoa) *errors.Aplicação
	Atualizar(ID, *entidades.Pessoa) *errors.Aplicação
	Pegar(ID) (*entidades.Pessoa, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Curso representa quais são as opereçãoes necessárias para salvar e
// alterar um Curso definitivamente.
type Curso interface {
	InserirMatérias(*[]entidades.CursoMatéria) *errors.Aplicação
	Inserir(*entidades.Curso) *errors.Aplicação
	AtualizarMatérias(*[]entidades.CursoMatéria) *errors.Aplicação
	Atualizar(ID, *entidades.Curso) *errors.Aplicação
	PegarMatérias(ID) (*[]entidades.CursoMatéria, *errors.Aplicação)
	Pegar(ID) (*entidades.Curso, *errors.Aplicação)
	DeletarMatérias(ID) *errors.Aplicação
	Deletar(ID) *errors.Aplicação
}

// Aluno representa quais são as opereçãoes necessárias para salvar e
// alterar uma Aluno definitivamente.
type Aluno interface {
	Inserir(*entidades.Aluno) *errors.Aplicação
	Atualizar(ID, *entidades.Aluno) *errors.Aplicação
	Pegar(ID) (*entidades.Aluno, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Professor representa quais são as opereçãoes necessárias para salvar e
// alterar uma Professor definitivamente.
type Professor interface {
	Inserir(*entidades.Professor) *errors.Aplicação
	Atualizar(ID, *entidades.Professor) *errors.Aplicação
	Pegar(ID) (*entidades.Professor, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Administrativo representa quais são as opereçãoes necessárias para salvar e
// alterar uma Administrativo definitivamente.
type Administrativo interface {
	Inserir(*entidades.Administrativo) *errors.Aplicação
	Atualizar(ID, *entidades.Administrativo) *errors.Aplicação
	Pegar(ID) (*entidades.Administrativo, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Matéria representa quais são as opereçãoes necessárias para salvar e
// alterar uma Matéria definitivamente.
type Matéria interface {
	Inserir(*entidades.Matéria) *errors.Aplicação
	Atualizar(ID, *entidades.Matéria) *errors.Aplicação
	Pegar(ID) (*entidades.Matéria, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Turma representa quais são as opereçãoes necessárias para salvar e
// alterar uma Turma definitivamente.
type Turma interface {
	Inserir(*entidades.Turma) *errors.Aplicação
	Atualizar(ID, *entidades.Turma) *errors.Aplicação
	Pegar(ID) (*entidades.Turma, *errors.Aplicação)
	Deletar(ID) *errors.Aplicação
}

// Data representa quais são as operações para modificar as entidades de uma
// forma definitiva.
type Data struct {
	Pessoa
	Curso
	Aluno
	Professor
	Administrativo
	Matéria
	Turma
}
