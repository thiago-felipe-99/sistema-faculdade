package errors

// Aplicação representa um erro na aplicação.
type Aplicação struct {
	Mensagem    string
	ErroInicial *Aplicação
	ErroExterno error
	Número      int
}

func (err *Aplicação) Error() string {
	return err.Mensagem
}

func (err *Aplicação) IsDefault(defaultError *Default) bool {
	return err.Número == defaultError.Número
}

// Default representa os erros padões da aplicação.
type Default struct {
	Message string
	Número  int
}

func (err *Default) Error() string {
	return err.Message
}

func New(err *Default, initial *Aplicação, system error) *Aplicação {
	return &Aplicação{
		Mensagem:    err.Message,
		Número:      err.Número,
		ErroInicial: initial,
		ErroExterno: system,
	}
}

func NewApplication(
	message string,
	número int,
	initialError *Aplicação,
	systemError error,
) *Aplicação {
	return &Aplicação{
		Mensagem:    message,
		Número:      número,
		ErroInicial: initialError,
		ErroExterno: systemError,
	}
}
