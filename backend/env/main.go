// Package env representa como a aplicação lida com o ambiente que roda.
package env

import (
	"os"
)

const (
	portaPadrãoBDAdministrativo = "3306"
	portaPadrãoBDMatéria        = "27017"
	hostPadrãoBDAdministrativo  = "db-administrativo"
	hostPadrãoBDMatéria         = "db-materia"
)

// Portas representa as variáveis de ambiente do tipo porta.
type Portas struct {
	BDAdministrativo string
	BDMateria        string
}

// Hosts representa os hosts necessários da aplicação.
type Hosts struct {
	BDAdministrativo string
	BDMateria        string
}

// VariáveisDeAmbiente são as variáveis do ambiente onde a aplicação roda.
type VariáveisDeAmbiente struct {
	Portas Portas
	Hosts  Hosts
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

	variáveis.Hosts.BDAdministrativo = os.Getenv("BD_ADMINISTRATIVO_HOST")
	if variáveis.Hosts.BDAdministrativo == "" {
		variáveis.Hosts.BDAdministrativo = hostPadrãoBDAdministrativo
	}

	variáveis.Hosts.BDMateria = os.Getenv("BD_MATERIA_HOST")
	if variáveis.Hosts.BDMateria == "" {
		variáveis.Hosts.BDMateria = hostPadrãoBDMatéria
	}

	return variáveis
}
