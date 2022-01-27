package logica

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type dataPessoaInvalida struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida) Inserir(pessoa *pessoa) erro {
	return erros.Novo(ErroCriarPessoa, nil, nil)
}

func (p *dataPessoaInvalida) Atualizar(id id, pessoa *pessoa) erro {
	return erros.Novo(ErroAtualizarPessoa, nil, nil)
}

func (p *dataPessoaInvalida) Pegar(id id) (*pessoa, erro) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida) PegarPorCPF(cpf cpf) (*pessoa, erro) {
	return p.data.PegarPorCPF(cpf)
}

func (p *dataPessoaInvalida) Deletar(id) erro {
	return erros.Novo(ErroDeletarPessoa, nil, nil)
}

type dataPessoaInvalida2 struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida2) Inserir(pessoa *pessoa) erro {
	return erros.Novo(ErroCriarPessoa, nil, nil)
}

func (p *dataPessoaInvalida2) Atualizar(id id, pessoa *pessoa) erro {
	return erros.Novo(ErroAtualizarPessoa, nil, nil)
}

func (p *dataPessoaInvalida2) Pegar(id id) (*pessoa, erro) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida2) PegarPorCPF(cpf cpf) (*pessoa, erro) {
	return nil, erros.Novo(ErroPegarPessoa, nil, nil)
}

func (p *dataPessoaInvalida2) Deletar(id) erro {
	return erros.Novo(ErroDeletarPessoa, nil, nil)
}

type dataMatériaInvalida struct {
	data data.Matéria
}

func (p *dataMatériaInvalida) Inserir(matéria *matéria) erro {
	return erros.Novo(ErroCriarMatéria, nil, nil)
}

func (p *dataMatériaInvalida) Atualizar(id id, matéria *matéria) erro {
	return erros.Novo(ErroAtualizarMatéria, nil, nil)
}

func (p *dataMatériaInvalida) PegarPréRequisitos(ids id) ([]id, erro) {
	return []id{}, erros.Novo(ErroAtualizarMatéria, nil, nil)
}

// func (p *dataMatériaInvalida) PegarMúltiplos(ids []id) ([]matéria, erro) {
// 	return []matéria{}, erros.Novo(ErroPegarMatéria, nil, nil)
// }

func (p *dataMatériaInvalida) PegarMúltiplos(ids []id) ([]matéria, erro) {
	return p.data.PegarMúltiplos(ids)
}

func (p *dataMatériaInvalida) Pegar(id id) (*matéria, erro) {
	return p.data.Pegar(id)
}

func (p *dataMatériaInvalida) Deletar(id) erro {
	return erros.Novo(ErroDeletarMatéria, nil, nil)
}

type dataMatériaInvalida2 struct {
	data data.Matéria
}

func (p *dataMatériaInvalida2) Inserir(matéria *matéria) erro {
	return erros.Novo(ErroCriarMatéria, nil, nil)
}

func (p *dataMatériaInvalida2) Atualizar(id id, matéria *matéria) erro {
	return erros.Novo(ErroAtualizarMatéria, nil, nil)
}

func (p *dataMatériaInvalida2) PegarPréRequisitos(ids id) ([]id, erro) {
	return []id{}, erros.Novo(ErroAtualizarMatéria, nil, nil)
}

func (p *dataMatériaInvalida2) PegarMúltiplos(ids []id) ([]matéria, erro) {
	return []matéria{}, erros.Novo(ErroPegarMatéria, nil, nil)
}

func (p *dataMatériaInvalida2) Pegar(id id) (*matéria, erro) {
	return p.data.Pegar(id)
}

func (p *dataMatériaInvalida2) Deletar(id) erro {
	return erros.Novo(ErroDeletarMatéria, nil, nil)
}

type dataCursoInvalido struct {
	data data.Curso
}

func (p *dataCursoInvalido) Inserir(curso *curso) erro {
	return erros.Novo(ErroCriarCurso, nil, nil)
}

func (p *dataCursoInvalido) Atualizar(id id, curso *curso) erro {
	return erros.Novo(ErroAtualizarCurso, nil, nil)
}

func (p *dataCursoInvalido) Pegar(id id) (*curso, erro) {
	return p.data.Pegar(id)
}

func (p *dataCursoInvalido) Deletar(id) erro {
	return erros.Novo(ErroDeletarCurso, nil, nil)
}
