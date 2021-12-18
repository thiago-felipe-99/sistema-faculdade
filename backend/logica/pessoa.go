package logica

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/erros"

	//nolint:revive,stylecheck
	. "thiagofelipe.com.br/sistema-faculdade/logica/erros"
)

type Pessoa struct {
	data data.Pessoa
}

func (p *Pessoa) ExisteCPF(cpf entidades.CPF) (bool, *erros.Aplicação) {
	pessoas, erro := p.data.PegarPorCPF(cpf)
	if erro != nil {
		return true, erros.Novo(ErroAoVerificarCPF, erro, nil)
	}

	return len(*pessoas) != 0, nil
}

func (p *Pessoa) Criar(
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
) (*entidades.Pessoa, *erros.Aplicação) {
	cpf, cpfVálido := entidades.ValidarCPF(cpf)
	if !cpfVálido {
		return nil, erros.Novo(ErroCPFInválido, nil, nil)
	}

	cpfExiste, erro := p.ExisteCPF(cpf)
	if erro != nil {
		return nil, erros.Novo(ErroCriarPessoa, erro, nil)
	}

	if cpfExiste {
		return nil, erros.Novo(ErroCPFExiste, nil, nil)
	}

	dataDeNascimento = entidades.RemoverHorário(dataDeNascimento.UTC())

	if dataDeNascimento.After(entidades.DataAtual()) {
		return nil, erros.Novo(ErroDataDeNascimentoInválido, nil, nil)
	}

	if !entidades.SenhaVálida(senha) {
		return nil, erros.Novo(ErroSenhaInválida, nil, nil)
	}

	pessoaNova := &entidades.Pessoa{
		ID:               entidades.NovoID(),
		Nome:             nome,
		CPF:              cpf,
		DataDeNascimento: dataDeNascimento,
		Senha:            entidades.GerarNovaSenha(senha),
	}

	erro = p.data.Inserir(pessoaNova)
	if erro != nil {
		return nil, erros.Novo(ErroCriarPessoa, erro, nil)
	}

	return pessoaNova, nil
}
