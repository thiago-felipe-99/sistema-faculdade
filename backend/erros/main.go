package erros

import "fmt"

// Aplicação representa um erro na aplicação.
type Aplicação struct {
	Mensagem    string
	ErroInicial *Aplicação
	ErroExterno error
	Código      string
}

func (erro *Aplicação) Traçado() string {
	mensagem := fmt.Sprintf("Erro Da Aplicação[%s]: %s", erro.Código, erro.Mensagem)

	if erro.ErroExterno != nil {
		mensagem += fmt.Sprintf("\nErro Externo: %s", erro.ErroExterno.Error())
	}

	if erro.ErroInicial != nil {
		mensagem += "\n" + erro.ErroInicial.Traçado()
	}

	return mensagem
}

func (erro *Aplicação) Error() string {
	return erro.Traçado()
}

func (erro *Aplicação) ÉPadrão(defaultError *Padrão) bool {
	return erro.Código == defaultError.Código
}

// Padrão representa os erros padrões da aplicação.
type Padrão struct {
	Mensagem string
	Código   string
}

func (erro *Padrão) Error() string {
	return fmt.Sprintf("Erro[%s]: %s", erro.Código, erro.Mensagem)
}

func Novo(err *Padrão, initial *Aplicação, system error) *Aplicação {
	return &Aplicação{
		Mensagem:    err.Mensagem,
		Código:      err.Código,
		ErroInicial: initial,
		ErroExterno: system,
	}
}

func ErroExterno(erro error) string {
	return fmt.Sprintf("Erro Externo: %s", erro.Error())
}
