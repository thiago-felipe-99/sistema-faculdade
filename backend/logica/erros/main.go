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
)

var (
	ErroAoVerificarCPF = &erros.Padrão{
		Mensagem: "Erro ao verificar o CPF",
		Código:   fmt.Sprintf("LÓGICA-[%d]", verificarCPF),
	}
	ErroCPFExiste = &erros.Padrão{
		Mensagem: "Já existe uma pessoa com esse CPF",
		Código:   fmt.Sprintf("LÓGICA-[%d]", cpfExiste),
	}
	ErroCPFInválido = &erros.Padrão{
		Mensagem: "CPF Inválido",
		Código:   fmt.Sprintf("LÓGICA-[%d]", cpfInválido),
	}
	ErroDataDeNascimentoInválido = &erros.Padrão{
		Mensagem: "Data de nascimento é inválido",
		Código:   fmt.Sprintf("LÓGICA-[%d]", dataDeNascimentoInválido),
	}
	ErroSenhaInválida = &erros.Padrão{
		Mensagem: "A senha não está dentro dos padrões de uma senha valida",
		Código:   fmt.Sprintf("LÓGICA-[%d]", senhaInválida),
	}
	ErroCriarPessoa = &erros.Padrão{
		Mensagem: "Erro Ao Criar A Pessoa",
		Código:   fmt.Sprintf("LÓGICA-[%d]", criarPessoa),
	}
)
