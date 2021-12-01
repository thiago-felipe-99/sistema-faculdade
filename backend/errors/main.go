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

func (err *Aplicação) IsDefault(defaultError *Padrão) bool {
	return err.Número == defaultError.Número
}

// Padrão representa os erros padões da aplicação.
type Padrão struct {
	Mensagem string
	Número   int
}

func (err *Padrão) Error() string {
	return err.Mensagem
}

func New(err *Padrão, initial *Aplicação, system error) *Aplicação {
	return &Aplicação{
		Mensagem:    err.Mensagem,
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
