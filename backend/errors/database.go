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
	CursoNãoEncontrada = &Padrão{
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
	CursoMatériasNãoEncontrada = &Padrão{
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
)
