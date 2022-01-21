package logica //nolint:testpackage

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type dataPessoaInvalida struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida) Inserir(pessoa *entidades.Pessoa) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida) Atualizar(
	id entidades.ID,
	pessoa *entidades.Pessoa,
) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida) Pegar(id entidades.ID) (
	*entidades.Pessoa,
	*erros.Aplicação,
) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida) PegarPorCPF(cpf entidades.CPF) (
	*entidades.Pessoa,
	*erros.Aplicação,
) {
	return p.data.PegarPorCPF(cpf)
}

func (p *dataPessoaInvalida) Deletar(entidades.ID) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

type dataPessoaInvalida2 struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida2) Inserir(pessoa *entidades.Pessoa) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Atualizar(
	id entidades.ID,
	pessoa *entidades.Pessoa,
) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Pegar(id entidades.ID) (
	*entidades.Pessoa,
	*erros.Aplicação,
) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida2) PegarPorCPF(cpf entidades.CPF) (
	*entidades.Pessoa,
	*erros.Aplicação,
) {
	return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Deletar(entidades.ID) *erros.Aplicação {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}
