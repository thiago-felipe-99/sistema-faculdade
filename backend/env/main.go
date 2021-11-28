package env

import (
	"os"
)

const portaPadrãoDBAdministrativo = "9000"

const portaPadrãoDBMatéria = "9001"

type Portas struct {
	BDAdministrativo string
	BDMateria        string
}

type VariáveisDeAmbiente struct {
	Portas Portas
}

func PegandoVariáveisDeAmbiente() (variáveis VariáveisDeAmbiente) {
	variáveis.Portas.BDAdministrativo = os.Getenv("DB_ADMINISTRATIVO_PORT")
	if variáveis.Portas.BDAdministrativo == "" {
		variáveis.Portas.BDAdministrativo = portaPadrãoDBAdministrativo
	}

	variáveis.Portas.BDMateria = os.Getenv("DB_MATERIA_PORT")
	if variáveis.Portas.BDMateria == "" {
		variáveis.Portas.BDMateria = portaPadrãoDBMatéria
	}

	return variáveis
}
