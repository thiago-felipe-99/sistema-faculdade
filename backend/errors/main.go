package errors

// Application representa um erro na aplicação.
type Application struct {
	Message      string
	InitialError *Application
	SystemError  error
}

func (err *Application) Error() string {
	return err.Message
}

func NewApplication(
	message string,
	initialError *Application,
	systemError error,
) *Application {
	return &Application{
		Message:      message,
		InitialError: initialError,
		SystemError:  systemError,
	}
}
