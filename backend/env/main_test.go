package env

import (
	"os"
	"reflect"
	"testing"
)

func TestPegandoVariáveisDeAmbiente(t *testing.T) {
	t.Run("SemVariáveisDeAmbiente", func(t *testing.T) {
		esperado := VariáveisDeAmbiente{
			Portas: Portas{
				BDAdministrativo: portaPadrãoBDAdministrativo,
				BDMateria:        portaPadrãoBDMatéria,
			},
		}

		recebido := PegandoVariáveisDeAmbiente()

		if !reflect.DeepEqual(esperado, recebido) {
			t.Fatalf("Esperava: %v\nChegou: %v", esperado, recebido)
		}
	})

	t.Run("ComVariáveisDeAmbiente", func(t *testing.T) {
		esperado := VariáveisDeAmbiente{
			Portas: Portas{
				BDAdministrativo: portaPadrãoBDAdministrativo + "0",
				BDMateria:        portaPadrãoBDMatéria + "0",
			},
		}

		erro := os.Setenv("BD_ADMINISTRATIVO_PORT", esperado.Portas.BDAdministrativo)
		if erro != nil {
			t.Fatalf("Não esperva nenhum erro, porém recebeu: %v", erro)
		}

		erro = os.Setenv("BD_MATERIA_PORT", esperado.Portas.BDMateria)
		if erro != nil {
			t.Fatalf("Não esperva nenhum erro, porém recebeu: %v", erro)
		}

		recebido := PegandoVariáveisDeAmbiente()

		if !reflect.DeepEqual(esperado, recebido) {
			t.Fatalf("Esperava: %v\nChegou: %v", esperado, recebido)
		}
	})
}
