package env

import (
	"os"
)

const (
	portaPadrãoBDAdministrativo = "9000"
	portaPadrãoBDMatéria        = "9001"
)

type Portas struct {
	BDAdministrativo string
	BDMateria        string
}

type VariáveisDeAmbiente struct {
	Portas Portas
}

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
