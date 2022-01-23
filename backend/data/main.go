package data

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type (
	erro           = *erros.Aplicação
	id             = entidades.ID
	cpf            = entidades.CPF
	pessoa         = entidades.Pessoa
	curso          = entidades.Curso
	cursomatéria   = entidades.CursoMatéria
	aluno          = entidades.Aluno
	professor      = entidades.Professor
	administrativo = entidades.Administrativo
	matéria        = entidades.Matéria
	turma          = entidades.Turma
)

// Pessoa representa quais são as opereçãoes necessárias para salvar e
// alterar uma pessoa definitivamente.
type Pessoa interface {
	Inserir(*pessoa) erro
	Atualizar(id, *pessoa) erro
	Pegar(id) (*pessoa, erro)
	PegarPorCPF(cpf) (*pessoa, erro)
	Deletar(id) erro
}

// Curso representa quais são as opereçãoes necessárias para salvar e
// alterar um Curso definitivamente.
type Curso interface {
	InserirMatérias(*[]cursomatéria) erro
	Inserir(*curso) erro
	AtualizarMatérias(*[]cursomatéria) erro
	Atualizar(id, *curso) erro
	PegarMatérias(id) (*[]cursomatéria, erro)
	Pegar(id) (*curso, erro)
	DeletarMatérias(id) erro
	Deletar(id) erro
}

// Aluno representa quais são as opereçãoes necessárias para salvar e
// alterar uma Aluno definitivamente.
type Aluno interface {
	Inserir(*aluno) erro
	Atualizar(id, *aluno) erro
	Pegar(id) (*aluno, erro)
	Deletar(id) erro
}

// Professor representa quais são as opereçãoes necessárias para salvar e
// alterar uma Professor definitivamente.
type Professor interface {
	Inserir(*professor) erro
	Atualizar(id, *professor) erro
	Pegar(id) (*professor, erro)
	Deletar(id) erro
}

// Administrativo representa quais são as opereçãoes necessárias para salvar e
// alterar uma Administrativo definitivamente.
type Administrativo interface {
	Inserir(*administrativo) erro
	Atualizar(id, *administrativo) erro
	Pegar(id) (*administrativo, erro)
	Deletar(id) erro
}

// Matéria representa quais são as opereçãoes necessárias para salvar e
// alterar uma Matéria definitivamente.
type Matéria interface {
	Inserir(*matéria) erro
	Atualizar(id, *matéria) erro
	Pegar(id) (*matéria, erro)
	ExisteIDs([]id) ([]id, bool, erro)
	Deletar(id) erro
}

// Turma representa quais são as opereçãoes necessárias para salvar e
// alterar uma Turma definitivamente.
type Turma interface {
	Inserir(*turma) erro
	Atualizar(id, *turma) erro
	Pegar(id) (*turma, erro)
	Deletar(id) erro
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
