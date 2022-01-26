package logica

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
	idsTamanho
	verificarIDs
	verificarID
	criarMatéria
	cargaHoráriaSemanalInválida
	créditosInválido
	préRequisitosNãoExiste
	matériaNãoEncontrada
	atualizarMatéria
	pegarMatéria
	deletarMatéria
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
	ErroIDsTamanho = criarErroPadrão(
		"Não foi passado nenhum ID",
		idsTamanho,
	)
	ErroVerificarID = criarErroPadrão(
		"Erro ao verificar se os ID existe",
		verificarID,
	)
	ErroVerificarIDs = criarErroPadrão(
		"Erro ao verificar se os IDs existe",
		verificarIDs,
	)
	ErroCriarMatéria = criarErroPadrão(
		"Erro ao criar a matéria",
		criarMatéria,
	)
	ErroCargaHoráriaMínima = criarErroPadrão(
		"A carga horária semanal mínima é de 1 hora",
		cargaHoráriaSemanalInválida,
	)
	ErroCréditosInválido = criarErroPadrão(
		"Créditos deve ser maior que 0",
		créditosInválido,
	)
	ErroPréRequisitosNãoExiste = criarErroPadrão(
		"Um dos pré requisitos não existe",
		préRequisitosNãoExiste,
	)
	ErroMatériaNãoEncontrada = criarErroPadrão(
		"Não foi encontrado a matéria na aplicação",
		matériaNãoEncontrada,
	)
	ErroAtualizarMatéria = criarErroPadrão(
		"Erro ao atualizar a matéria",
		atualizarMatéria,
	)
	ErroPegarMatéria = criarErroPadrão(
		"Erro ao pegar a matéria",
		pegarMatéria,
	)
	ErroDeletarMatéria = criarErroPadrão(
		"Erro ao deletar a matéria",
		deletarMatéria,
	)
)
