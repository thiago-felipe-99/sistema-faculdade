package erros

import (
	"fmt"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

const (
	configurarBD = iota + 1
	inserirPessoa
	atualizarPessoa
	pessoaNãoEncontrada
	pegarPessoa
	pegarPessoaCPF
	DeletarPessoa
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

var (
	ErroConfigurarBD = &erros.Padrão{
		Mensagem: "Erro ao configurar o banco de dados",
		Código:   fmt.Sprintf("DATA-[%d]", configurarBD),
	}
	ErroInserirPessoa = &erros.Padrão{
		Mensagem: "Erro ao inserir a pessoa",
		Código:   fmt.Sprintf("DATA-[%d]", inserirPessoa),
	}
	ErroAtualizarPessoa = &erros.Padrão{
		Mensagem: "Erro ao atualizar a pessoa",
		Código:   fmt.Sprintf("DATA-[%d]", atualizarPessoa),
	}
	ErroPessoaNãoEncontrada = &erros.Padrão{
		Mensagem: "Pessoa não encontrada",
		Código:   fmt.Sprintf("DATA-[%d]", pessoaNãoEncontrada),
	}
	ErroPegarPessoa = &erros.Padrão{
		Mensagem: "Erro ao pegar a pessoa",
		Código:   fmt.Sprintf("DATA-[%d]", pegarPessoa),
	}
	ErroPegarPessoaPorCPF = &erros.Padrão{
		Mensagem: "Erro ao pegar a pessoa pelo CPF",
		Código:   fmt.Sprintf("DATA-[%d]", pegarPessoaCPF),
	}
	ErroDeletarPessoa = &erros.Padrão{
		Mensagem: "Erro ao deletar a pessoa",
		Código:   fmt.Sprintf("DATA-[%d]", DeletarPessoa),
	}
	ErroInserirCurso = &erros.Padrão{
		Mensagem: "Erro ao inserir o curso",
		Código:   fmt.Sprintf("DATA-[%d]", inserirCurso),
	}
	ErroAtualizarCurso = &erros.Padrão{
		Mensagem: "Erro ao atualizar o curso",
		Código:   fmt.Sprintf("DATA-[%d]", atualizarCurso),
	}
	ErroCursoNãoEncontrado = &erros.Padrão{
		Mensagem: "Curso não encontrada",
		Código:   fmt.Sprintf("DATA-[%d]", cursoNãoEncontrado),
	}
	ErroPegarCurso = &erros.Padrão{
		Mensagem: "Erro ao pegar o curso",
		Código:   fmt.Sprintf("DATA-[%d]", pegarCurso),
	}
	ErroDeletarCurso = &erros.Padrão{
		Mensagem: "Erro ao deletar o curso",
		Código:   fmt.Sprintf("DATA-[%d]", deletarCurso),
	}
	ErroInserirCursoMatériasTamanhoMínimo = &erros.Padrão{
		//nolint:lll
		Mensagem: "Erro ao inserir as matérias do curso, tem que ter no mínimo uma matéra para inserir",
		Código:   fmt.Sprintf("DATA-[%d]", inserirCursoMatériasTamanhoMínimo),
	}
	ErroInserirCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao inserir as matérias do curso",
		Código:   fmt.Sprintf("DATA-[%d]", inserirCursoMatérias),
	}
	ErroAtualizarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao atualizar as matérias do curso",
		Código:   fmt.Sprintf("DATA-[%d]", atualizarCursoMatérias),
	}
	ErroCursoMatériasNãoEncontrado = &erros.Padrão{
		Mensagem: "Matérias do curso não encontradas",
		Código:   fmt.Sprintf("DATA-[%d]", cursoMatériasNãoEncontrado),
	}
	ErroPegarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao pegar as matérias do curso",
		Código:   fmt.Sprintf("DATA-[%d]", pegarCursoMatérias),
	}
	ErroDeletarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao deletar as matérias do curso",
		Código:   fmt.Sprintf("DATA-[%d]", deletarCursoMatérias),
	}
	ErroInserirAluno = &erros.Padrão{
		Mensagem: "Erro ao inserir o aluno",
		Código:   fmt.Sprintf("DATA-[%d]", inserirAluno),
	}
	ErroAtualizarAluno = &erros.Padrão{
		Mensagem: "Erro ao atualizar o aluno",
		Código:   fmt.Sprintf("DATA-[%d]", atualizarAluno),
	}
	ErroAlunoNãoEncontrado = &erros.Padrão{
		Mensagem: "Aluno não encontrada",
		Código:   fmt.Sprintf("DATA-[%d]", alunoNãoEncontrado),
	}
	ErroPegarAluno = &erros.Padrão{
		Mensagem: "Erro ao pegar o aluno",
		Código:   fmt.Sprintf("DATA-[%d]", pegarAluno),
	}
	ErroDeletarAluno = &erros.Padrão{
		Mensagem: "Erro ao deletar o aluno",
		Código:   fmt.Sprintf("DATA-[%d]", deletarAluno),
	}
	ErroInserirAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao inserir as turmas do aluno",
		Código:   fmt.Sprintf("DATA-[%d]", inserirAlunoTurma),
	}
	ErroAtualizarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao atualizar as turmas do aluno",
		Código:   fmt.Sprintf("DATA-[%d]", atualizarAlunoTurma),
	}
	ErroAlunoTurmaNãoEncontrado = &erros.Padrão{
		Mensagem: "Turmas do aluno não encontradas",
		Código:   fmt.Sprintf("DATA-[%d]", alunoTurmaNãoEncontrado),
	}
	ErroPegarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao pegar as turmas do aluno",
		Código:   fmt.Sprintf("DATA-[%d]", pegarAlunoTurma),
	}
	ErroDeletarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao deletar as turmas do aluno",
		Código:   fmt.Sprintf("DATA-[%d]", deletarAlunoTurma),
	}
)
