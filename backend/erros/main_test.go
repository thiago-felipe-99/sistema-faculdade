package erros

import (
	"fmt"
	"testing"
)

func TestAplicação(t *testing.T) {
	erroPadrão1 := Padrão{"Mensagem1", "Código1"}
	erroPadrão2 := Padrão{"Mensagem2", "Código2"}
	erroInicial := &Aplicação{
		Mensagem:    "MensagemProfunda",
		ErroInicial: nil,
		ErroExterno: &erroPadrão2,
		Código:      "CódigoProdundo",
	}

	erro := Novo(&erroPadrão1, erroInicial, &erroPadrão2)

	t.Run("Traçado", func(t *testing.T) {
		mensagem := fmt.Sprintf(
			"Erro Da Aplicação[%s]: %s\n\t%s\n\t%s",
			erroPadrão1.Código, erroPadrão1.Mensagem,
			ErroExterno(&erroPadrão2),
			erroInicial.Traçado(),
		)

		if mensagem != erro.Traçado() {
			t.Fatalf("Esperava: %s\nChegou: %s", mensagem, erro.Traçado())
		}

	})

	t.Run("Error", func(t *testing.T) {
		if erro.Error() != erro.Traçado() {
			t.Fatalf("Esperava: %s\nChegou: %s", erro.Traçado(), erro.Error())
		}
	})

	t.Run("ÉPadrão", func(t *testing.T) {
		if !erro.ÉPadrão(&erroPadrão1) {
			t.Fatalf("Esperava: %s\nChegou: %s", erroPadrão1.Código, erro.Código)
		}
	})
}

func TestPadrão(t *testing.T) {
	erro := Padrão{"Mensagem", "Código"}
	mensagem := "Erro[Código]: Mensagem"

	if erro.Error() != mensagem {
		t.Fatalf("Esperava: %s\nChegou: %s", mensagem, erro.Error())
	}
}

func TestNovo(t *testing.T) {
	erro := Padrão{"Mensagem", "Código"}

	erroEsperado := &Aplicação{
		Mensagem:    "Mensagem",
		Código:      "Código",
		ErroInicial: nil,
		ErroExterno: &erro,
	}

	erroRecebido := Novo(&erro, nil, &erro)

	if erroEsperado.Error() != erroRecebido.Error() {
		t.Fatalf("Esperava: %s\nChegou: %s", erroEsperado, erroRecebido)
	}
}

func TestErroExterno(t *testing.T) {
	erro := Padrão{"Mensagem", "Código"}
	mensagem := "Erro Externo: Erro[Código]: Mensagem"

	if ErroExterno(&erro) != mensagem {
		t.Fatalf("Esperava: %s\nChegou: %s", mensagem, erro.Error())
	}
}
