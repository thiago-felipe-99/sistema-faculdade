package erros

import (
	"fmt"

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
)

func criarErroPadrão(mensagem string, códigoNúmero int) *erros.Padrão {
	return &erros.Padrão{
		Mensagem: mensagem,
		Código:   fmt.Sprintf("LÓGICA-[%d]", códigoNúmero),
	}
}

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
)
