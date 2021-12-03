package erros

const (
	ConfigurarBDNúmero = iota + 1
	InserirPessoaNúmero
	AtualizarPessoaNúmero
	PessoaNãoEncontradaNúmero
	PegarPessoaNúmero
	DeletarPessoaNúmero
	InserirCursoNúmero
	AtualizarCursoNúmero
	CursoNãoEncontradoNúmero
	PegarCursoNúmero
	DeletarCursoNúmero
	InserirCursoMatériasNúmero
	InserirCursoMatériasTamanhoMínimoNúmero
	AtualizarCursoMatériasNúmero
	CursoMatériasNãoEncontradoNúmero
	PegarCursoMatériasNúmero
	DeletarCursoMatériasNúmero
	InserirAlunoNúmero
	AtualizarAlunoNúmero
	AlunoNãoEncontradoNúmero
	PegarAlunoNúmero
	DeletarAlunoNúmero
	InserirAlunoTurmaNúmero
	AtualizarAlunoTurmaNúmero
	AlunoTurmaNãoEncontradoNúmero
	PegarAlunoTurmaNúmero
	DeletarAlunoTurmaNúmero
)

var (
	ConfigurarBD = &Padrão{
		Mensagem: "Erro ao configurar o banco de dados",
		Número:   ConfigurarBDNúmero,
	}
	InserirPessoa = &Padrão{
		Mensagem: "Erro ao inserir a pessoa",
		Número:   InserirPessoaNúmero,
	}
	AtualizarPessoa = &Padrão{
		Mensagem: "Erro ao atualizar a pessoa",
		Número:   AtualizarPessoaNúmero,
	}
	PessoaNãoEncontrada = &Padrão{
		Mensagem: "Pessoa não encontrada",
		Número:   PessoaNãoEncontradaNúmero,
	}
	PegarPessoa = &Padrão{
		Mensagem: "Erro ao pegar a pessoa",
		Número:   PegarPessoaNúmero,
	}
	DeletarPessoa = &Padrão{
		Mensagem: "Erro ao deletar a pessoa",
		Número:   DeletarPessoaNúmero,
	}
	InserirCurso = &Padrão{
		Mensagem: "Erro ao inserir o curso",
		Número:   InserirCursoNúmero,
	}
	AtualizarCurso = &Padrão{
		Mensagem: "Erro ao atualizar o curso",
		Número:   AtualizarCursoNúmero,
	}
	CursoNãoEncontrado = &Padrão{
		Mensagem: "Curso não encontrada",
		Número:   CursoNãoEncontradoNúmero,
	}
	PegarCurso = &Padrão{
		Mensagem: "Erro ao pegar o curso",
		Número:   PegarCursoNúmero,
	}
	DeletarCurso = &Padrão{
		Mensagem: "Erro ao deletar o curso",
		Número:   DeletarCursoNúmero,
	}
	InserirCursoMatériasTamanhoMínimo = &Padrão{
		Mensagem: "Erro ao inserir as matérias do curso, tem que ter no mínimo uma matéra para inserir",
		Número:   InserirCursoMatériasTamanhoMínimoNúmero,
	}
	InserirCursoMatérias = &Padrão{
		Mensagem: "Erro ao inserir as matérias do curso",
		Número:   InserirCursoMatériasNúmero,
	}
	AtualizarCursoMatérias = &Padrão{
		Mensagem: "Erro ao atualizar as matérias do curso",
		Número:   AtualizarCursoMatériasNúmero,
	}
	CursoMatériasNãoEncontrado = &Padrão{
		Mensagem: "Matérias do curso não encontradas",
		Número:   CursoMatériasNãoEncontradoNúmero,
	}
	PegarCursoMatérias = &Padrão{
		Mensagem: "Erro ao pegar as matérias do curso",
		Número:   PegarCursoMatériasNúmero,
	}
	DeletarCursoMatérias = &Padrão{
		Mensagem: "Erro ao deletar as matérias do curso",
		Número:   DeletarCursoMatériasNúmero,
	}
	InserirAluno = &Padrão{
		Mensagem: "Erro ao inserir o aluno",
		Número:   InserirAlunoNúmero,
	}
	AtualizarAluno = &Padrão{
		Mensagem: "Erro ao atualizar o aluno",
		Número:   AtualizarAlunoNúmero,
	}
	AlunoNãoEncontrado = &Padrão{
		Mensagem: "Aluno não encontrada",
		Número:   AlunoNãoEncontradoNúmero,
	}
	PegarAluno = &Padrão{
		Mensagem: "Erro ao pegar o aluno",
		Número:   PegarAlunoNúmero,
	}
	DeletarAluno = &Padrão{
		Mensagem: "Erro ao deletar o aluno",
		Número:   DeletarAlunoNúmero,
	}
	InserirAlunoTurma = &Padrão{
		Mensagem: "Erro ao inserir as turmas do aluno",
		Número:   InserirAlunoTurmaNúmero,
	}
	AtualizarAlunoTurma = &Padrão{
		Mensagem: "Erro ao atualizar as turmas do aluno",
		Número:   AtualizarAlunoTurmaNúmero,
	}
	AlunoTurmaNãoEncontrado = &Padrão{
		Mensagem: "Turmas do aluno não encontradas",
		Número:   AlunoTurmaNãoEncontradoNúmero,
	}
	PegarAlunoTurma = &Padrão{
		Mensagem: "Erro ao pegar as turmas do aluno",
		Número:   PegarAlunoTurmaNúmero,
	}
	DeletarAlunoTurma = &Padrão{
		Mensagem: "Erro ao deletar as turmas do aluno",
		Número:   DeletarAlunoTurmaNúmero,
	}
)
