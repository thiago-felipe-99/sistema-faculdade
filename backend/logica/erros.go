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
	verificarID
	verificarIDs
	verificarPréRequisitosCiclos
	criarMatéria
	cargaHoráriaSemanalInválida
	créditosInválido
	préRequisitosNãoExiste
	matériaNãoEncontrada
	atualizarMatéria
	préRequisitosCiclo
	pegarMatéria
	deletarMatéria
	criarCurso
	dataDeInícioMaior
	atualizarCurso
	cursoNãoEncontrado
	pegarCurso
	deletarCurso
)

var criarErroPadrão = erros.NovoPadrãoFunc("LÓGICA") //nolint:gochecknoglobals

// Possíveis erros do pacote.
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
	ErroVerificarPréRequisitosCiclos = criarErroPadrão(
		"Erro ao verificar se á ciclos nos pré-requisitos",
		verificarPréRequisitosCiclos,
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
	ErroPréRequisitosCiclo = criarErroPadrão(
		"Ao um ciclo entre os pré-requisitos",
		préRequisitosCiclo,
	)
	ErroPegarMatéria = criarErroPadrão(
		"Erro ao pegar a matéria",
		pegarMatéria,
	)
	ErroDeletarMatéria = criarErroPadrão(
		"Erro ao deletar a matéria",
		deletarMatéria,
	)
	ErroCriarCurso = criarErroPadrão(
		"Erro ao criar a curso",
		criarCurso,
	)
	ErroDataDeInícioMaior = criarErroPadrão(
		"Data de início é maior que a data final",
		dataDeInícioMaior,
	)
	ErroCursoNãoEncontrado = criarErroPadrão(
		"Não foi encontrado o curso na aplicação",
		cursoNãoEncontrado,
	)
	ErroAtualizarCurso = criarErroPadrão(
		"Erro ao atualizar a curso",
		atualizarCurso,
	)
	ErroPegarCurso = criarErroPadrão(
		"Erro ao pegar o curso",
		pegarCurso,
	)
	ErroDeletarCurso = criarErroPadrão(
		"Erro ao deletar o curso",
		deletarCurso,
	)
)
