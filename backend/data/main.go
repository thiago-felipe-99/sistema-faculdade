package data

import (
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/erros"
)

type ID = entidades.ID

// Pessoa representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type Pessoa interface {
	Inserir(*entidades.Pessoa) *erros.Aplicação
	Atualizar(ID, *entidades.Pessoa) *erros.Aplicação
	Pegar(ID) (*entidades.Pessoa, *erros.Aplicação)
	BuscarPorCPF(cpf entidades.CPF) (*[]entidades.Pessoa, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
}

// Curso representa quais são as opereçãoes necessárias para salvar e
// alterar um Curso definitivamente.
type Curso interface {
	InserirMatérias(*[]entidades.CursoMatéria) *erros.Aplicação
	Inserir(*entidades.Curso) *erros.Aplicação
	AtualizarMatérias(*[]entidades.CursoMatéria) *erros.Aplicação
	Atualizar(ID, *entidades.Curso) *erros.Aplicação
	PegarMatérias(ID) (*[]entidades.CursoMatéria, *erros.Aplicação)
	Pegar(ID) (*entidades.Curso, *erros.Aplicação)
	DeletarMatérias(ID) *erros.Aplicação
	Deletar(ID) *erros.Aplicação
}

// Aluno representa quais são as opereçãoes necessárias para salvar e
// alterar uma Aluno definitivamente.
type Aluno interface {
	Inserir(*entidades.Aluno) *erros.Aplicação
	Atualizar(ID, *entidades.Aluno) *erros.Aplicação
	Pegar(ID) (*entidades.Aluno, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
}

// Professor representa quais são as opereçãoes necessárias para salvar e
// alterar uma Professor definitivamente.
type Professor interface {
	Inserir(*entidades.Professor) *erros.Aplicação
	Atualizar(ID, *entidades.Professor) *erros.Aplicação
	Pegar(ID) (*entidades.Professor, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
}

// Administrativo representa quais são as opereçãoes necessárias para salvar e
// alterar uma Administrativo definitivamente.
type Administrativo interface {
	Inserir(*entidades.Administrativo) *erros.Aplicação
	Atualizar(ID, *entidades.Administrativo) *erros.Aplicação
	Pegar(ID) (*entidades.Administrativo, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
}

// Matéria representa quais são as opereçãoes necessárias para salvar e
// alterar uma Matéria definitivamente.
type Matéria interface {
	Inserir(*entidades.Matéria) *erros.Aplicação
	Atualizar(ID, *entidades.Matéria) *erros.Aplicação
	Pegar(ID) (*entidades.Matéria, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
}

// Turma representa quais são as opereçãoes necessárias para salvar e
// alterar uma Turma definitivamente.
type Turma interface {
	Inserir(*entidades.Turma) *erros.Aplicação
	Atualizar(ID, *entidades.Turma) *erros.Aplicação
	Pegar(ID) (*entidades.Turma, *erros.Aplicação)
	Deletar(ID) *erros.Aplicação
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
