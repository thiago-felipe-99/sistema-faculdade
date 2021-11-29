package errors

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
)

var (
	ConfigurarBD = &Default{
		Message: "Erro ao configurar o banco de dados",
		Número:  ConfigurarBDNúmero,
	}
	InserirPessoa = &Default{
		Message: "Erro ao inserir a pessoa",
		Número:  InserirPessoaNúmero,
	}
	AtualizarPessoa = &Default{
		Message: "Erro ao atualizar a pessoa",
		Número:  AtualizarPessoaNúmero,
	}
	PessoaNãoEncontrada = &Default{
		Message: "Pessoa não encontrada",
		Número:  PessoaNãoEncontradaNúmero,
	}
	PegarPessoa = &Default{
		Message: "Erro ao pegar a pessoa",
		Número:  PegarPessoaNúmero,
	}
	DeletarPessoa = &Default{
		Message: "Erro ao deletar a pessoa",
		Número:  DeletarPessoaNúmero,
	}
	InserirCurso = &Default{
		Message: "Erro ao inserir o curso",
		Número:  InserirCursoNúmero,
	}
	AtualizarCurso = &Default{
		Message: "Erro ao atualizar o curso",
		Número:  AtualizarCursoNúmero,
	}
	CursoNãoEncontrada = &Default{
		Message: "Curso não encontrada",
		Número:  CursoNãoEncontradoNúmero,
	}
	PegarCurso = &Default{
		Message: "Erro ao pegar o curso",
		Número:  PegarCursoNúmero,
	}
	DeletarCurso = &Default{
		Message: "Erro ao deletar o curso",
		Número:  DeletarCursoNúmero,
	}
	InserirCursoMatériasTamanhoMínimo = &Default{
		Message: "Erro ao inserir as matérias do curso, tem que ter no mínimo uma matéra para inserir",
		Número:  InserirCursoMatériasTamanhoMínimoNúmero,
	}
	InserirCursoMatérias = &Default{
		Message: "Erro ao inserir as matérias do curso",
		Número:  InserirCursoMatériasNúmero,
	}
	AtualizarCursoMatérias = &Default{
		Message: "Erro ao atualizar as matérias do curso",
		Número:  AtualizarCursoMatériasNúmero,
	}
	CursoMatériasNãoEncontrada = &Default{
		Message: "Matérias do curso não encontradas",
		Número:  CursoMatériasNãoEncontradoNúmero,
	}
	PegarCursoMatérias = &Default{
		Message: "Erro ao pegar as matérias do curso",
		Número:  PegarCursoMatériasNúmero,
	}
	DeletarCursoMatérias = &Default{
		Message: "Erro ao deletar as matérias do curso",
		Número:  DeletarCursoMatériasNúmero,
	}
)
