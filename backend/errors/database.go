package errors

const (
	ConfigurarBDNúmero = iota + 1
	InserirPessoaNúmero
	AtualizarPessoaNúmero
	PessoaNãoEncontradaNúmero
	PegarPessoaNúmero
	DeletarPessoaNúmero
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
)
