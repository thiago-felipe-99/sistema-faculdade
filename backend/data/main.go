package data

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// Pessoa representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type Pessoa interface {
	Inserir(*entidades.Pessoa) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Pessoa) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Pessoa, *erros.Aplicação)
	PegarPorCPF(entidades.CPF) (*entidades.Pessoa, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
}

// Curso representa quais são as opereçãoes necessárias para salvar e
// alterar um Curso definitivamente.
type Curso interface {
	InserirMatérias(*[]entidades.CursoMatéria) *erros.Aplicação
	Inserir(*entidades.Curso) *erros.Aplicação
	AtualizarMatérias(*[]entidades.CursoMatéria) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Curso) *erros.Aplicação
	PegarMatérias(entidades.ID) (*[]entidades.CursoMatéria, *erros.Aplicação)
	Pegar(entidades.ID) (*entidades.Curso, *erros.Aplicação)
	DeletarMatérias(entidades.ID) *erros.Aplicação
	Deletar(entidades.ID) *erros.Aplicação
}

// Aluno representa quais são as opereçãoes necessárias para salvar e
// alterar uma Aluno definitivamente.
type Aluno interface {
	Inserir(*entidades.Aluno) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Aluno) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Aluno, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
}

// Professor representa quais são as opereçãoes necessárias para salvar e
// alterar uma Professor definitivamente.
type Professor interface {
	Inserir(*entidades.Professor) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Professor) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Professor, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
}

// Administrativo representa quais são as opereçãoes necessárias para salvar e
// alterar uma Administrativo definitivamente.
type Administrativo interface {
	Inserir(*entidades.Administrativo) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Administrativo) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Administrativo, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
}

// Matéria representa quais são as opereçãoes necessárias para salvar e
// alterar uma Matéria definitivamente.
type Matéria interface {
	Inserir(*entidades.Matéria) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Matéria) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Matéria, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
}

// Turma representa quais são as opereçãoes necessárias para salvar e
// alterar uma Turma definitivamente.
type Turma interface {
	Inserir(*entidades.Turma) *erros.Aplicação
	Atualizar(entidades.ID, *entidades.Turma) *erros.Aplicação
	Pegar(entidades.ID) (*entidades.Turma, *erros.Aplicação)
	Deletar(entidades.ID) *erros.Aplicação
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
