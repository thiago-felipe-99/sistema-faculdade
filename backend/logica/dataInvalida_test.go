package logica

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

type dataPessoaInvalida struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida) Inserir(pessoa *pessoa) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida) Atualizar(id id, pessoa *pessoa) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida) Pegar(id id) (*pessoa, erro) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida) PegarPorCPF(cpf cpf) (*pessoa, erro) {
	return p.data.PegarPorCPF(cpf)
}

func (p *dataPessoaInvalida) Deletar(id) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

type dataPessoaInvalida2 struct {
	data data.Pessoa
}

func (p *dataPessoaInvalida2) Inserir(pessoa *pessoa) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Atualizar(id id, pessoa *pessoa) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Pegar(id id) (*pessoa, erro) {
	return p.data.Pegar(id)
}

func (p *dataPessoaInvalida2) PegarPorCPF(cpf cpf) (*pessoa, erro) {
	return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}

func (p *dataPessoaInvalida2) Deletar(id) erro {
	return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
}
