//nolint: paralleltest
package aleatorio_test

import (
	"testing"
	"unicode/utf8"

	. "thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func TestNúmero(t *testing.T) {
	const tamanho uint = 1000

	t.Run("OKAY", func(t *testing.T) {
		máximo := Número(tamanho)

		if tamanho <= máximo {
			t.Fatalf(
				"O número deveria ser maior que 0 e menor que %d, porém chegou %d.",
				tamanho,
				máximo,
			)
		}
	})

	t.Run("0", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if r != ErroTamanhoInválido.Error() {
					t.Fatalf("\nEsperava: %v\nChegou  : %v", ErroTamanhoInválido.Error(), r)
				}
			} else {
				t.Fatalf("Esperar ocorrer um Panic")
			}
		}()

		Número(0)
	})
}

func TestPalavra(t *testing.T) {
	const tamanho uint = 10

	t.Run("OKAY", func(t *testing.T) {
		palavra := Palavra(tamanho)
		palavraTamanho := utf8.RuneCountInString(palavra)

		if palavraTamanho != int(tamanho) {
			t.Fatalf("Tamanho esperado: %d, chegou: %d", tamanho, palavraTamanho)
		}
	})

	t.Run("0", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if r != ErroTamanhoInválido.Error() {
					t.Fatalf("\nEsperava: %v\nChegou  : %v", ErroTamanhoInválido.Error(), r)
				}
			} else {
				t.Fatalf("Esperava ocorrer um Panic")
			}
		}()

		Palavra(0)
	})
}

func TestCPF(t *testing.T) {
	cpf := CPF() //nolint: ifshort
	if _, válido := entidades.ValidarCPF(cpf); !válido {
		t.Fatalf("Esperava um CPF válido, chegou: %s", cpf)
	}
}

func TestBytes(t *testing.T) {
	const tamanho uint32 = 50

	if bytes := Bytes(tamanho); len(bytes) != int(tamanho) {
		t.Fatalf("Tamanho esperado: %d, chegou: %d", tamanho, len(bytes))
	}
}

func TestSenha(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
		s := Senha()

		gerenciador := &entidades.Senha{}

		if !gerenciador.ÉVálida(s) {
			t.Fatalf("Esperava uma senha válida, chegou: %s", s)
		}
	})
}
