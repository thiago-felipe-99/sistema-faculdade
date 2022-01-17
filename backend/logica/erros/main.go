package erros

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

const (
	verificarCPF = iota + 1
	cpfExiste
	cpfInválido
	dataDeNascimentoInválido
	senhaInválida
	criarPessoa
	pessoaNãoEncontrada
	pegarPessoa
	verificarSenha
	atualizarPessoa
	deletarPessoa
)

var criarErroPadrão = erros.NovoPadrãoFunc("LÓGICA") //nolint:gochecknoglobals

var (
	ErroAoVerificarCPF = criarErroPadrão(
		"Erro ao verificar o CPF",
		verificarCPF,
	)
	ErroCPFExiste = criarErroPadrão(
		"Já existe uma pessoa com esse CPF",
		cpfExiste,
	)
	ErroCPFInválido = criarErroPadrão(
		"CPF Inválido",
		cpfInválido,
	)
	ErroDataDeNascimentoInválido = criarErroPadrão(
		"Data de nascimento é inválido",
		dataDeNascimentoInválido,
	)
	ErroSenhaInválida = criarErroPadrão(
		"A senha não está dentro dos padrões de uma senha valida",
		senhaInválida,
	)
	ErroCriarPessoa = criarErroPadrão(
		"Erro Ao Criar A Pessoa",
		criarPessoa,
	)
	ErroPessoaNãoEncontrada = criarErroPadrão(
		"Pessoa não econtrada",
		pessoaNãoEncontrada,
	)
	ErroPegarPessoa = criarErroPadrão(
		"Erro ao pegar a pessoa",
		pegarPessoa,
	)
	ErroVerificarSenha = criarErroPadrão(
		"Erro ao verificar a senha",
		verificarSenha,
	)
	ErroAtualizarPessoa = criarErroPadrão(
		"Erro ao atualizar a pessoa",
		atualizarPessoa,
	)
	ErroDeletarPessoa = criarErroPadrão(
		"Erro ao deletar a pessoa",
		deletarPessoa,
	)
)
