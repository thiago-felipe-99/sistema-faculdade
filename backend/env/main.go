// Package env representa como a aplicação lida com o ambiente que roda.
package env

import (
	"os"
)

const (
	portaPadrãoBDAdministrativo = "9000"
	portaPadrãoBDMatéria        = "9001"
)

// Portas representa as variáveis de ambiente do tipo porta.
type Portas struct {
	BDAdministrativo string
	BDMateria        string
}

// VariáveisDeAmbiente são as variáveis do ambiente onde a aplicação roda.
type VariáveisDeAmbiente struct {
	Portas Portas
}

// PegandoVariáveisDeAmbiente retorna as variáveis de ambiente que a aplicação
// precisa.
func PegandoVariáveisDeAmbiente() (variáveis VariáveisDeAmbiente) {
	variáveis.Portas.BDAdministrativo = os.Getenv("BD_ADMINISTRATIVO_PORT")
	if variáveis.Portas.BDAdministrativo == "" {
		variáveis.Portas.BDAdministrativo = portaPadrãoBDAdministrativo
	}

	variáveis.Portas.BDMateria = os.Getenv("BD_MATERIA_PORT")
	if variáveis.Portas.BDMateria == "" {
		variáveis.Portas.BDMateria = portaPadrãoBDMatéria
	}

	return variáveis
}
