package data

import (
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

const (
	configurarBD = iota + 1
	inserirPessoa
	atualizarPessoa
	pessoaNãoEncontrada
	pegarPessoa
	pegarPessoaCPF
	deletarPessoa
	inserirCurso
	atualizarCurso
	cursoNãoEncontrado
	pegarCurso
	deletarCurso
	inserirCursoMatérias
	inserirCursoMatériasTamanhoMínimo
	atualizarCursoMatérias
	cursoMatériasNãoEncontrado
	pegarCursoMatérias
	deletarCursoMatérias
	inserirAluno
	atualizarAluno
	alunoNãoEncontrado
	pegarAluno
	deletarAluno
	inserirAlunoTurma
	atualizarAlunoTurma
	alunoTurmaNãoEncontrado
	pegarAlunoTurma
	deletarAlunoTurma
)

var criarErroPadrão = erros.NovoPadrãoFunc("LÓGICA") //nolint:gochecknoglobals

var (
	ErroConfigurarBD = criarErroPadrão(
		"Erro ao configurar o banco de dados",
		configurarBD,
	)
	ErroInserirPessoa = criarErroPadrão(
		"Erro ao inserir a pessoa",
		inserirPessoa,
	)
	ErroAtualizarPessoa = criarErroPadrão(
		"Erro ao atualizar a pessoa",
		atualizarPessoa,
	)
	ErroPessoaNãoEncontrada = criarErroPadrão(
		"Pessoa não encontrada",
		pessoaNãoEncontrada,
	)
	ErroPegarPessoa = criarErroPadrão(
		"Erro ao pegar a pessoa",
		pegarPessoa,
	)
	ErroPegarPessoaPorCPF = criarErroPadrão(
		"Erro ao pegar a pessoa pelo CPF",
		pegarPessoaCPF,
	)
	ErroDeletarPessoa = criarErroPadrão(
		"Erro ao deletar a pessoa",
		deletarPessoa,
	)
	ErroInserirCurso = criarErroPadrão(
		"Erro ao inserir o curso",
		inserirCurso,
	)
	ErroAtualizarCurso = criarErroPadrão(
		"Erro ao atualizar o curso",
		atualizarCurso,
	)
	ErroCursoNãoEncontrado = criarErroPadrão(
		"Curso não encontrada",
		cursoNãoEncontrado,
	)
	ErroPegarCurso = criarErroPadrão(
		"Erro ao pegar o curso",
		pegarCurso,
	)
	ErroDeletarCurso = criarErroPadrão(
		"Erro ao deletar o curso",
		deletarCurso,
	)
	ErroInserirCursoMatériasTamanhoMínimo = criarErroPadrão(
		"Erro ao inserir as matérias do curso, tem que ter no mínimo uma matéra para inserir", //nolint:lll
		inserirCursoMatériasTamanhoMínimo,
	)
	ErroInserirCursoMatérias = criarErroPadrão(
		"Erro ao inserir as matérias do curso",
		inserirCursoMatérias,
	)
	ErroAtualizarCursoMatérias = criarErroPadrão(
		"Erro ao atualizar as matérias do curso",
		atualizarCursoMatérias,
	)
	ErroCursoMatériasNãoEncontrado = criarErroPadrão(
		"Matérias do curso não encontradas",
		cursoMatériasNãoEncontrado,
	)
	ErroPegarCursoMatérias = criarErroPadrão(
		"Erro ao pegar as matérias do curso",
		pegarCursoMatérias,
	)
	ErroDeletarCursoMatérias = criarErroPadrão(
		"Erro ao deletar as matérias do curso",
		deletarCursoMatérias,
	)
	ErroInserirAluno = criarErroPadrão(
		"Erro ao inserir o aluno",
		inserirAluno,
	)
	ErroAtualizarAluno = criarErroPadrão(
		"Erro ao atualizar o aluno",
		atualizarAluno,
	)
	ErroAlunoNãoEncontrado = criarErroPadrão(
		"Aluno não encontrada",
		alunoNãoEncontrado,
	)
	ErroPegarAluno = criarErroPadrão(
		"Erro ao pegar o aluno",
		pegarAluno,
	)
	ErroDeletarAluno = criarErroPadrão(
		"Erro ao deletar o aluno",
		deletarAluno,
	)
	ErroInserirAlunoTurma = criarErroPadrão(
		"Erro ao inserir as turmas do aluno",
		inserirAlunoTurma,
	)
	ErroAtualizarAlunoTurma = criarErroPadrão(
		"Erro ao atualizar as turmas do aluno",
		atualizarAlunoTurma,
	)
	ErroAlunoTurmaNãoEncontrado = criarErroPadrão(
		"Turmas do aluno não encontradas",
		alunoTurmaNãoEncontrado,
	)
	ErroPegarAlunoTurma = criarErroPadrão(
		"Erro ao pegar as turmas do aluno",
		pegarAlunoTurma,
	)
	ErroDeletarAlunoTurma = criarErroPadrão(
		"Erro ao deletar as turmas do aluno",
		deletarAlunoTurma,
	)
)
