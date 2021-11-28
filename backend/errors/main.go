package errors

// Application representa um erro na aplicação.
type Application struct {
	Message      string
	InitialError *Application
	SystemError  error
	Número       int
}

func (err *Application) Error() string {
	return err.Message
}

func (err *Application) IsDefault(defaultError *Default) bool {
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

func New(err *Default, initial *Application, system error) *Application {
	return &Application{
		Message:      err.Message,
		Número:       err.Número,
		InitialError: initial,
		SystemError:  system,
	}
}

func NewApplication(
	message string,
	número int,
	initialError *Application,
	systemError error,
) *Application {
	return &Application{
		Message:      message,
		Número:       número,
		InitialError: initialError,
		SystemError:  systemError,
	}
}
