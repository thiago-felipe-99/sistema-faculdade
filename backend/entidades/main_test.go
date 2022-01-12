package entidades

import (
	"fmt"
	"testing"
)

func TestParseCPF(t *testing.T) {
	var testes = []struct {
		cpf      string
		válido   bool
		cpfParse string
	}{
		{"12345678901", true, "12345678901"},
		{"123.456.789-01", true, "12345678901"},
		{"123.45678901", false, "00000000000"},
		{"123.456.78901", false, "00000000000"},
		{"123.456.789-010", false, "00000000000"},
		{"0123.456.789-01", false, "00000000000"},
		{"123456789090..-", false, "00000000000"},
	}
	for _, teste := range testes {
		nomeTeste := fmt.Sprintf("%s", teste.cpf)
		t.Run(nomeTeste, func(t *testing.T) {
			cpf, válido := parseCPF(teste.cpf)
			if válido != teste.válido {
				t.Errorf("Queria %t, chegou %t", teste.válido, válido)
			}
			if cpf != teste.cpfParse {
				t.Errorf("Queria %s, chegou %s", teste.cpfParse, cpf)
			}
		})
	}
}

func TestVerificarDígitosCPF(t *testing.T) {
	var testes = []struct {
		cpf    string
		válido bool
	}{
		{"123.45678901", false},
		{"123.456.78901", false},
		{"123.456.789-010", false},
		{"0123.456.789-01", false},
		{"12345678909..-", false},
		{"55688304014", true},
		{"556.883.040-14", true},
		{"55688304013", false},
		{"55688304024", false},
		{"24020435049", true},
		{"240.204.350-49", true},
		{"14020435049", false},
		{"24020435149", false},
	}

	for _, teste := range testes {
		nomeTeste := fmt.Sprintf("%s", teste.cpf)
		t.Run(nomeTeste, func(t *testing.T) {
			válido := verificarDígitoCPF(teste.cpf)
			if válido != teste.válido {
				t.Errorf("Queria %t, chegou %t", teste.válido, válido)
			}
		})
	}
}

func TestValidarCPF(t *testing.T) {
	var testes = []struct {
		cpf      string
		válido   bool
		cpfParse string
	}{
		{"12345678901", false, "00000000000"},
		{"123.456.789-01", false, "00000000000"},
		{"123.45678901", false, "00000000000"},
		{"123.456.78901", false, "00000000000"},
		{"123.456.789-010", false, "00000000000"},
		{"0123.456.789-01", false, "00000000000"},
		{"12345678909..-", false, "00000000000"},
		{"12345678909", true, "12345678909"},
		{"123.456.789-09", true, "12345678909"},
		{"123.45678909", false, "00000000000"},
		{"123.456.78909", false, "00000000000"},
		{"123.456.789-090", false, "00000000000"},
		{"0123.456.789-09", false, "00000000000"},
		{"123456789090..-", false, "00000000000"},
	}
	for _, teste := range testes {
		nomeTeste := fmt.Sprintf("%s", teste.cpf)
		t.Run(nomeTeste, func(t *testing.T) {
			cpf, válido := ValidarCPF(teste.cpf)
			if válido != teste.válido {
				t.Errorf("Queria %t, chegou %t", teste.válido, válido)
			}
			if cpf != teste.cpfParse {
				t.Errorf("Queria %s, chegou %s", teste.cpfParse, cpf)
			}
		})
	}
}
