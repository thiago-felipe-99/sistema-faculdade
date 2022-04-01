package logica

import (
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// Pessoa representa operações que se possa fazer com a entidade Pessoa.
type Pessoa struct {
	Data data.Pessoa
}

// existeMatéria verifica se a matéria existeMatéria na aplicação.
func (lógica *Pessoa) existe(id id) (bool, erro) {
	_, erro := lógica.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			return false, nil
		}

		return false, erros.Novo(ErroVerificarID, erro, nil)
	}

	return true, nil
}

// existeCPF procura se já existe uma pessoa com esse CPF na aplicação.
func (lógica *Pessoa) existeCPF(cpf cpf) (bool, erro) {
	_, erro := lógica.Data.PegarPorCPF(cpf)
	if erro != nil {
		if erro.ÉPadrão(data.ErroPessoaNãoEncontrada) {
			return false, nil
		}

		return true, erros.Novo(ErroAoVerificarCPF, erro, nil)
	}

	return true, nil
}

// Criar adiciona uma pessoa na aplicação.
func (lógica *Pessoa) Criar(
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
) (*pessoa, erro) {
	cpf, cpfVálido := entidades.ValidarCPF(cpf)
	if !cpfVálido {
		return nil, erros.Novo(ErroCPFInválido, nil, nil)
	}

	cpfExiste, erro := lógica.existeCPF(cpf)
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

	gerenciadorSenha := entidades.GerenciadorSenhaPadrão()
	if !gerenciadorSenha.ÉVálida(senha) {
		return nil, erros.Novo(ErroSenhaInválida, nil, nil)
	}

	pessoaNova := &pessoa{
		ID:               entidades.NovoID(),
		Nome:             nome,
		CPF:              cpf,
		DataDeNascimento: dataDeNascimento,
		Senha:            gerenciadorSenha.GerarHash(senha),
	}

	erro = lógica.Data.Inserir(pessoaNova)
	if erro != nil {
		return nil, erros.Novo(ErroCriarPessoa, erro, nil)
	}

	return pessoaNova, nil
}

// Pegar retorna uma pessoa já criada na aplicação.
func (lógica *Pessoa) Pegar(id id) (*pessoa, erro) {
	pessoa, erro := lógica.Data.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(data.ErroPessoaNãoEncontrada) {
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroPegarPessoa, erro, nil)
	}

	return pessoa, nil
}

// VerificarSenha verifica se a senha fornecida é igual a senha da Pessoa na
// aplicação.
func (lógica *Pessoa) VerificarSenha(senha string, id id) (bool, erro) {
	pessoa, erro := lógica.Pegar(id)
	if erro != nil {
		if erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			return false, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return false, erros.Novo(ErroVerificarSenha, erro, nil)
	}

	gerenciadorSenha := entidades.GerenciadorSenhaPadrão()

	igual, erro := gerenciadorSenha.ÉIgual(senha, pessoa.Senha)
	if erro != nil {
		return false, erros.Novo(ErroVerificarSenha, erro, nil)
	}

	return igual, nil
}

// Atualizar atualiza os dados de uma pessoa na aplicação.
func (lógica *Pessoa) Atualizar(
	idPessoa id,
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
) (*pessoa, erro) {
	pessoaSalva, erro := lógica.Pegar(idPessoa)
	if erro != nil {
		if erro.ÉPadrão(ErroPessoaNãoEncontrada) {
			return nil, erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
		}

		return nil, erros.Novo(ErroAtualizarPessoa, erro, nil)
	}

	cpf, cpfVálido := entidades.ValidarCPF(cpf)
	if !cpfVálido {
		return nil, erros.Novo(ErroCPFInválido, nil, nil)
	}

	if cpf != pessoaSalva.CPF {
		cpfExiste, erro := lógica.existeCPF(cpf)
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

	gerenciadorSenha := entidades.GerenciadorSenhaPadrão()
	if !gerenciadorSenha.ÉVálida(senha) {
		return nil, erros.Novo(ErroSenhaInválida, nil, nil)
	}

	pessoaNova := &pessoa{
		ID:               idPessoa,
		Nome:             nome,
		CPF:              cpf,
		DataDeNascimento: dataDeNascimento,
		Senha:            gerenciadorSenha.GerarHash(senha),
	}

	erro = lógica.Data.Atualizar(idPessoa, pessoaNova)
	if erro != nil {
		return nil, erros.Novo(ErroAtualizarPessoa, erro, nil)
	}

	return pessoaNova, nil
}

// Deletar remove uma pessoa da aplicação.
func (lógica *Pessoa) Deletar(idPessoa id) erro {
	existe, erro := lógica.existe(idPessoa)
	if erro != nil {
		return erros.Novo(ErroDeletarPessoa, erro, nil)
	}

	if !existe {
		return erros.Novo(ErroPessoaNãoEncontrada, nil, nil)
	}

	erro = lógica.Data.Deletar(idPessoa)
	if erro != nil {
		return erros.Novo(ErroDeletarPessoa, erro, nil)
	}

	return nil
}
