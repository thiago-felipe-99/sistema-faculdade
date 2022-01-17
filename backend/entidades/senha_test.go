package entidades

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"reflect"
	"testing"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

func TestVerificarSenha(t *testing.T) {
	gerenciadorSenha := GerenciadorSenhaPadrão()
	senhaPlana := aleatorio.Senha()
	hash := gerenciadorSenha.GerarHash(senhaPlana)

	t.Run("SemErro", func(t *testing.T) {
		testes := []struct {
			senha string
			igual bool
		}{
			{senhaPlana, true},
			{senhaPlana + "a", false},
			{senhaPlana[:len(senhaPlana)-1] + "a", false},
			{senhaPlana + senhaPlana, false},
			{"", false},
		}

		for _, teste := range testes {
			t.Run(teste.senha, func(t *testing.T) {
				igual, erro := gerenciadorSenha.ÉIgual(teste.senha, hash)
				if erro != nil {
					t.Fatalf("Erro inesperado aconteceu: %v", erro)
				}

				if igual != teste.igual {
					t.Fatalf("Esperava: %v, Chegou: %v", teste.igual, igual)
				}
			})
		}
	})

	t.Run("ComErro", func(t *testing.T) {
		t.Run("Base64Inválida", func(t *testing.T) {
			_, erro := gerenciadorSenha.ÉIgual(senhaPlana, "ÇçÇçÇçÇçÇçÇç")

			var erroExterno base64.CorruptInputError = 0
			erroEsperadoInicial := erros.Novo(ErroDecodificarBase64, nil, erroExterno)
			erroEsperado := erros.Novo(ErroVerificarSenhaHash, erroEsperadoInicial, nil)

			if erroEsperado.Error() != erro.Error() {
				t.Fatalf("Esperava: %v\nChegou: %v", ErroVerificarSenhaHash, erro)
			}
		})

		t.Run("HashTamanhoInválido", func(t *testing.T) {
			_, erro := gerenciadorSenha.ÉIgual(senhaPlana, "aAaA")

			erroEsperadoInicial := erros.Novo(ErroDesencriptarAESNonceSize, nil, nil)
			erroEsperado := erros.Novo(ErroVerificarSenhaHash, erroEsperadoInicial, nil)

			if erroEsperado.Error() != erro.Error() {
				t.Fatalf("Esperava: %v\nChegou: %v", ErroVerificarSenhaHash, erro)
			}
		})
	})

	t.Run("Argon2idErro", func(t *testing.T) {
		erroInicial := erros.Novo(ErroDecodificarHashInválido, nil, nil)
		erroEsperado := erros.Novo(ErroVerificarSenhaHash, erroInicial, nil)

		_, erro := verificarSenhaHashArgon2id(senhaPlana, "a")
		if erro.Error() != erroEsperado.Error() {
			t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, erro)
		}
	})

}

func TestEncriptarAES(t *testing.T) {
	senha := []byte("senha")
	chave := []byte("1616161616161616")

	t.Run("ChaveInválida", func(t *testing.T) {
		defer func() {
			r := recover()
			var erroEsperado aes.KeySizeError = 1
			if !reflect.DeepEqual(r, erros.ErroExterno(erroEsperado)) {
				t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, r)
			}
		}()

		encriptarAES(senha, []byte("1"), 10)
	})

	t.Run("NonceSizeInválido", func(t *testing.T) {
		defer func() {
			r := recover()
			var erroEsperado = "Erro Externo: cipher: the nonce can't have zero length, or the security of the key will be immediately compromised"
			if r != erroEsperado {
				t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, r)
			}
		}()

		encriptarAES(senha, chave, 0)
	})
}

func TestDesencriptarAES(t *testing.T) {
	senha := []byte("senha")
	chave := []byte("1616161616161616")

	t.Run("ChaveInválida", func(t *testing.T) {
		var erroExterno aes.KeySizeError = 1
		erroEsperado := erros.Novo(ErroDesencriptarAES, nil, erroExterno)

		_, erro := desencriptarAES(senha, []byte("1"), 10)
		if erroEsperado.Error() != erro.Error() {
			t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, erro)
		}
	})

	t.Run("NonceSizeInválido", func(t *testing.T) {
		_, erro := desencriptarAES(senha, chave, 0)
		if erro == nil || !erro.ÉPadrão(ErroDesencriptarAES) {
			t.Fatalf("Esperava: %v\nChegou: %v", ErroDesencriptarAES, erro)
		}

		externo := "cipher: the nonce can't have zero length, or the security of the key will be immediately compromised"

		if erro.ErroExterno == nil || erro.ErroExterno.Error() != externo {
			t.Fatalf("Esperava erroInicial: %v\nChegou: %v", externo, erro.ErroExterno)
		}
	})

	t.Run("SenhaMenorNonceSize", func(t *testing.T) {
		erroEsperado := erros.Novo(ErroDesencriptarAESNonceSize, nil, nil)

		_, erro := desencriptarAES(senha, []byte("1616161616161616"), 16)
		if erroEsperado.Error() != erro.Error() {
			t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, erro)
		}
	})

	t.Run("HashInválido", func(t *testing.T) {
		erroInicial := errors.New("cipher: message authentication failed")
		erroEsperado := erros.Novo(ErroDesencriptarAES, nil, erroInicial)

		_, erro := desencriptarAES(senha, []byte("1616161616161616"), 1)
		if erroEsperado.Error() != erro.Error() {
			t.Fatalf("Esperava: %v\nChegou: %v", erroEsperado, erro)
		}
	})

}

func TestBase64DecodificarArgon2id(t *testing.T) {
	t.Run("OKAY", func(t *testing.T) {
		testes := []struct {
			senha, sal []byte
		}{
			{[]byte("Senha"), []byte("Senha")},
			{[]byte("Senha"), []byte("Teste")},
			{[]byte("password"), []byte("Teste")},
		}
		for _, teste := range testes {
			t.Run(string(teste.senha)+"-"+string(teste.sal), func(t *testing.T) {
				config := Argon2Config{
					memory:      64 * 1024,
					iterations:  3,
					parallelism: 2,
					saltLength:  uint32(len(teste.sal)),
					keyLength:   uint32(len(teste.senha)),
				}

				hash := base64CodificarArgon2id(teste.senha, teste.sal, config)

				senha, sal, configRecebido, erro := base64DecodificarArgon2id(hash)
				if erro != nil {
					t.Fatalf("Não esperava erro, chegou: %v", erro)
				}
				if !bytes.Equal(senha, teste.senha) {
					t.Fatalf("Esperava senha: %x\nChegou: %x", teste.senha, senha)
				}
				if !bytes.Equal(sal, teste.sal) {
					t.Fatalf("Esperava sal: %x\nChegou: %x", teste.sal, sal)
				}
				if !reflect.DeepEqual(config, *configRecebido) {
					t.Fatalf("Esperava config: %v\nChegou: %v", config, configRecebido)
				}
			})
		}
	})

	t.Run("Erros", func(t *testing.T) {
		testes := []struct {
			hash string
			erro *erros.Padrão
		}{
			{"hash", ErroDecodificarHashInválido},
			{"$hash$hash$hash$hash$hash", ErroDecodificarHashInválido},
			{"$argon2id$v19$hash$hash$hash", ErroDecodificarArgon2id},
			{"$argon2id$v=199$hash$hash$hash", ErroDecodificarArgon2idVersão},
			{"$argon2id$v=19$hash$hash$hash", ErroDecodificarArgon2id},
			{"$argon2id$v=19$m=16,t=2,p=1$Çash$hash", ErroDecodificarArgon2id},
			{"$argon2id$v=19$m=16,t=2,p=1$hash$çash", ErroDecodificarArgon2id},
		}

		for _, teste := range testes {
			t.Run(teste.hash, func(t *testing.T) {
				_, _, _, erro := base64DecodificarArgon2id(teste.hash)
				if erro == nil || !erro.ÉPadrão(teste.erro) {
					t.Fatalf("Esperava: %v\nChegou: %v", teste.erro, erro)
				}
			})
		}
	})
}

func TestSenhaVálida(t *testing.T) {
	testes := []struct {
		senha  string
		válida bool
	}{
		{"aA0-", false},
		{
			`aA0-0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`,
			false,
		},
		{"aAaAaAaAaAaA-", false},
		{"0123456789A-", false},
		{"0123456789a-", false},
		{"0123456789aA", false},
		{"0123456789aA-", true},
		{"0123456789çÇ-", true},
	}

	senha := GerenciadorSenhaPadrão()

	for _, teste := range testes {
		t.Run(teste.senha, func(t *testing.T) {
			válida := senha.ÉVálida(teste.senha)
			if válida != teste.válida {
				t.Fatalf("Esperava: %t, chegou: %t", teste.válida, válida)
			}
		})
	}
}

func TestGerenciadorSenhaPadrão(t *testing.T) {
	esperado := &Senha{
		argon2: Argon2Config{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
		},
		chave:     []byte("meMudeMeMudeMeMudeMeMudeMeMudeMe"),
		nonceSize: 16,
	}
	recebido := GerenciadorSenhaPadrão()

	if !reflect.DeepEqual(esperado, recebido) {
		t.Fatalf("Esperava: %v\nRecebou: %v", esperado, recebido)
	}
}
