package erros

import "thiagofelipe.com.br/sistema-faculdade/erros"

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
		Número:   verificarCPF,
	}
	ErroCPFExiste = &erros.Padrão{
		Mensagem: "Já existe uma pessoa com esse CPF",
		Número:   cpfExiste,
	}
	ErroCPFInválido = &erros.Padrão{
		Mensagem: "CPF Inválido",
		Número:   cpfInválido,
	}
	ErroDataDeNascimentoInválido = &erros.Padrão{
		Mensagem: "Data de nascimento é inválido",
		Número:   dataDeNascimentoInválido,
	}
	ErroSenhaInválida = &erros.Padrão{
		Mensagem: "A senha não está dentro dos padrões de uma senha valida",
		Número:   senhaInválida,
	}
	ErroCriarPessoa = &erros.Padrão{
		Mensagem: "Erro Ao Criar A Pessoa",
		Número:   criarPessoa,
	}
)
