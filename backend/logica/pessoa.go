package logica

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	dataErros "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"

	//nolint:revive,stylecheck
	. "thiagofelipe.com.br/sistema-faculdade-backend/logica/erros"
)

type Pessoa struct {
	data data.Pessoa
}

func (lógica *Pessoa) ExisteCPF(cpf entidades.CPF) (bool, *erros.Aplicação) {
	_, erro := lógica.data.PegarPorCPF(cpf)
	if erro != nil {
		if erro.ÉPadrão(dataErros.ErroPessoaNãoEncontrada) {
			return false, nil
		}

		return true, erros.Novo(ErroAoVerificarCPF, erro, nil)
	}

	return true, nil
}

func (lógica *Pessoa) Criar(
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
) (*entidades.Pessoa, *erros.Aplicação) {
	cpf, cpfVálido := entidades.ValidarCPF(cpf)
	if !cpfVálido {
		return nil, erros.Novo(ErroCPFInválido, nil, nil)
	}

	cpfExiste, erro := lógica.ExisteCPF(cpf)
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

	erro = lógica.data.Inserir(pessoaNova)
	if erro != nil {
		return nil, erros.Novo(ErroCriarPessoa, erro, nil)
	}

	return pessoaNova, nil
}

func (lógica *Pessoa) Pegar(id entidades.ID) (*entidades.Pessoa, *erros.Aplicação) {
	pessoa, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(dataErros.ErroPessoaNãoEncontrada) {
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroPegarPessoa, erro, nil)
	}

	return pessoa, nil
}

func (lógica *Pessoa) VerificarSenha(
	senha string,
	id entidades.ID,
) (bool, *erros.Aplicação) {
	pessoa, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(dataErros.ErroPessoaNãoEncontrada) {
			return false, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return false, erros.Novo(ErroPegarPessoa, erro, nil)
	}

	return entidades.VerificarSenha(senha, pessoa.Senha), nil
}

func (lógica *Pessoa) Atualizar(
	id entidades.ID,
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
) (*entidades.Pessoa, *erros.Aplicação) {
	pessoa, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(dataErros.ErroPessoaNãoEncontrada) {
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroAtualizarPessoa, erro, nil)
	}

	cpf, cpfVálido := entidades.ValidarCPF(cpf)
	if !cpfVálido {
		return nil, erros.Novo(ErroCPFInválido, nil, nil)
	}

	if cpf != pessoa.CPF {
		cpfExiste, erro := lógica.ExisteCPF(cpf)
		if erro != nil {
			return nil, erros.Novo(ErroAtualizarPessoa, erro, nil)
		}

		if cpfExiste {
			return nil, erros.Novo(ErroCPFExiste, nil, nil)
		}
	}

	dataDeNascimento = entidades.RemoverHorário(dataDeNascimento.UTC())

	if dataDeNascimento.After(entidades.DataAtual()) {
		return nil, erros.Novo(ErroDataDeNascimentoInválido, nil, nil)
	}

	if !entidades.SenhaVálida(senha) {
		return nil, erros.Novo(ErroSenhaInválida, nil, nil)
	}

	pessoaNova := &entidades.Pessoa{
		ID:               id,
		Nome:             nome,
		CPF:              cpf,
		DataDeNascimento: dataDeNascimento,
		Senha:            entidades.GerarNovaSenha(senha),
	}

	erro = lógica.data.Atualizar(id, pessoaNova)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarPessoa, erro, nil)
	}

	return pessoaNova, nil
}

func (lógica *Pessoa) Deletar(id entidades.ID) *erros.Aplicação {
	_, erro := lógica.data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(dataErros.ErroPessoaNãoEncontrada) {
			return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return erros.Novo(ErroDeletarPessoa, erro, nil)
	}

	erro = lógica.data.Deletar(id)
	if erro != nil {
		return erros.Novo(ErroDeletarPessoa, erro, nil)
	}

	return nil
}
