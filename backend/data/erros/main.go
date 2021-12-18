package erros

import "thiagofelipe.com.br/sistema-faculdade/erros"

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
		Número:   configurarBD,
	}
	ErroInserirPessoa = &erros.Padrão{
		Mensagem: "Erro ao inserir a pessoa",
		Número:   inserirPessoa,
	}
	ErroAtualizarPessoa = &erros.Padrão{
		Mensagem: "Erro ao atualizar a pessoa",
		Número:   atualizarPessoa,
	}
	ErroPessoaNãoEncontrada = &erros.Padrão{
		Mensagem: "Pessoa não encontrada",
		Número:   pessoaNãoEncontrada,
	}
	ErroPegarPessoa = &erros.Padrão{
		Mensagem: "Erro ao pegar a pessoa",
		Número:   pegarPessoa,
	}
	ErroPegarPessoaPorCPF = &erros.Padrão{
		Mensagem: "Erro ao pegar a pessoa pelo CPF",
		Número:   pegarPessoaCPF,
	}
	ErroDeletarPessoa = &erros.Padrão{
		Mensagem: "Erro ao deletar a pessoa",
		Número:   DeletarPessoa,
	}
	ErroInserirCurso = &erros.Padrão{
		Mensagem: "Erro ao inserir o curso",
		Número:   inserirCurso,
	}
	ErroAtualizarCurso = &erros.Padrão{
		Mensagem: "Erro ao atualizar o curso",
		Número:   atualizarCurso,
	}
	ErroCursoNãoEncontrado = &erros.Padrão{
		Mensagem: "Curso não encontrada",
		Número:   cursoNãoEncontrado,
	}
	ErroPegarCurso = &erros.Padrão{
		Mensagem: "Erro ao pegar o curso",
		Número:   pegarCurso,
	}
	ErroDeletarCurso = &erros.Padrão{
		Mensagem: "Erro ao deletar o curso",
		Número:   deletarCurso,
	}
	ErroInserirCursoMatériasTamanhoMínimo = &erros.Padrão{
		Mensagem: "Erro ao inserir as matérias do curso, tem que ter no mínimo uma matéra para inserir",
		Número:   inserirCursoMatériasTamanhoMínimo,
	}
	ErroInserirCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao inserir as matérias do curso",
		Número:   inserirCursoMatérias,
	}
	ErroAtualizarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao atualizar as matérias do curso",
		Número:   atualizarCursoMatérias,
	}
	ErroCursoMatériasNãoEncontrado = &erros.Padrão{
		Mensagem: "Matérias do curso não encontradas",
		Número:   cursoMatériasNãoEncontrado,
	}
	ErroPegarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao pegar as matérias do curso",
		Número:   pegarCursoMatérias,
	}
	ErroDeletarCursoMatérias = &erros.Padrão{
		Mensagem: "Erro ao deletar as matérias do curso",
		Número:   deletarCursoMatérias,
	}
	ErroInserirAluno = &erros.Padrão{
		Mensagem: "Erro ao inserir o aluno",
		Número:   inserirAluno,
	}
	ErroAtualizarAluno = &erros.Padrão{
		Mensagem: "Erro ao atualizar o aluno",
		Número:   atualizarAluno,
	}
	ErroAlunoNãoEncontrado = &erros.Padrão{
		Mensagem: "Aluno não encontrada",
		Número:   alunoNãoEncontrado,
	}
	ErroPegarAluno = &erros.Padrão{
		Mensagem: "Erro ao pegar o aluno",
		Número:   pegarAluno,
	}
	ErroDeletarAluno = &erros.Padrão{
		Mensagem: "Erro ao deletar o aluno",
		Número:   deletarAluno,
	}
	ErroInserirAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao inserir as turmas do aluno",
		Número:   inserirAlunoTurma,
	}
	ErroAtualizarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao atualizar as turmas do aluno",
		Número:   atualizarAlunoTurma,
	}
	ErroAlunoTurmaNãoEncontrado = &erros.Padrão{
		Mensagem: "Turmas do aluno não encontradas",
		Número:   alunoTurmaNãoEncontrado,
	}
	ErroPegarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao pegar as turmas do aluno",
		Número:   pegarAlunoTurma,
	}
	ErroDeletarAlunoTurma = &erros.Padrão{
		Mensagem: "Erro ao deletar as turmas do aluno",
		Número:   deletarAlunoTurma,
	}
)
