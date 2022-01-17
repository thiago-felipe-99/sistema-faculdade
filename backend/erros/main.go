package erros

import (
	"fmt"
)

// Aplicação representa um erro na aplicação.
type Aplicação struct {
	Mensagem    string
	ErroInicial *Aplicação
	ErroExterno error
	Código      string
}

// Traçado retorna todos os erros que ocorreram na aplicação.
func (erro *Aplicação) Traçado() string {
	mensagem := fmt.Sprintf("Erro Da Aplicação[%s]: %s", erro.Código, erro.Mensagem)

	if erro.ErroExterno != nil {
		mensagem += "\n\t" + ErroExterno(erro.ErroExterno)
	}

	if erro.ErroInicial != nil {
		mensagem += "\n\t" + erro.ErroInicial.Traçado()
	}

	return mensagem
}

// Error é um método para adequar a interface err da biblioteca padrão (stdlb).
func (erro *Aplicação) Error() string {
	return erro.Traçado()
}

// ÉPadrão verifica se um erro vem de um certo erro do tipo Padrão.
func (erro *Aplicação) ÉPadrão(defaultError *Padrão) bool {
	return erro.Código == defaultError.Código
}

// Padrão representa os erros padrões da aplicação.
type Padrão struct {
	Mensagem string
	Código   string
}

// Error é um método para adequar a interface err da biblioteca padrão (stdlb).
func (erro *Padrão) Error() string {
	return fmt.Sprintf("Erro %s: %s", erro.Código, erro.Mensagem)
}

// NovoPadrãoFunc retorna uma função para criar erros padrões.
func NovoPadrãoFunc(nomePacote string) func(mensagem string, código int) *Padrão {
	return func(mensagem string, código int) *Padrão {
		return &Padrão{
			Mensagem: mensagem,
			Código:   fmt.Sprintf("%s-[%d]", nomePacote, código),
		}
	}
}

// Novo criar um novo erro na aplicação.
func Novo(err *Padrão, initial *Aplicação, system error) *Aplicação {
	return &Aplicação{
		Mensagem:    err.Mensagem,
		Código:      err.Código,
		ErroInicial: initial,
		ErroExterno: system,
	}
}

// ErroExterno encapsula um erro externo a aplicação.
func ErroExterno(erro error) string {
	return fmt.Sprintf("Erro Externo: %s", erro.Error())
}
