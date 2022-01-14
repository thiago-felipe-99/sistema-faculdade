package entidades

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	base64STLB "encoding/base64"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"
	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// Hash representa uma hash na aplicação.
type Hash = string

// Argon2Config representa as configurações usada no algoritmo Argon2.
type Argon2Config struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// Senha gerencia como as senhas são tratadas na aplicação.
type Senha struct {
	argon2    Argon2Config
	chave     []byte
	nonceSize uint
}

// NovaSenha criar um gerenciador de senhas.
func NovaSenha(chave string, argon2Config Argon2Config, nonceSize uint) *Senha {
	return &Senha{
		argon2:    argon2Config,
		chave:     []byte(chave),
		nonceSize: nonceSize,
	}
}

// ÉIgual verifica se um uma senhaPlana é igual o hash.
func (senha *Senha) ÉIgual(senhaPlana string, hash Hash) bool {
	return senhaPlana == hash
}

// ÉVálida verifica se a senha cumpri os requisitos de uma senha forte.
func (senha *Senha) ÉVálida(senhaPlana string) bool {
	return true
}

// GerarHash retorna a senhaPlana hasheada no algoritmo padrão.
func (senha *Senha) GerarHash(senhaPlana string) Hash {
	hashSHA512 := gerarHashSHA3_512([]byte(senhaPlana))

	hashArgon2id, saltArgon2id := gerarHashArgon2id(hashSHA512, senha.argon2)

	hashB64 := base64EncodeArgon2id(hashArgon2id, saltArgon2id, senha.argon2)

	hashAES := encriptarAES([]byte(hashB64), senha.chave, senha.nonceSize)

	hash := base64Encode(hashAES)

	return hash
}

// gerarHashSHA3_512 gera o hash de uma senha plana pelo algoritmo SHA3-512 e
// retorna o hash em um slice de bytes.
func gerarHashSHA3_512(senhaPlana []byte) []byte {
	sha3_512 := sha3.New512()

	_, err := sha3_512.Write(senhaPlana)
	if err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	hash := sha3_512.Sum(nil)

	return hash
}

// gerarHashArgon2id gera o hash de uma senha plana pelo algoritmo argon2id, e
// retona o hash e o sal usado para gerar ela.
func gerarHashArgon2id(senhaPlana []byte, config Argon2Config) (hash, sal []byte) {
	sal = aleatorio.Bytes(config.saltLength)

	hash = argon2.IDKey(
		senhaPlana,
		sal,
		config.iterations,
		config.memory,
		config.parallelism,
		config.keyLength,
	)

	return hash, sal
}

// base64Encode retorna o slice de bytes codificado na base64.
func base64Encode(bytes []byte) string {
	return base64STLB.RawStdEncoding.EncodeToString(bytes)
}

// base64EncodeArgon2id retorna o hash de um algoritmo argon2id codificado na
// base64.
func base64EncodeArgon2id(hash, sal []byte, config Argon2Config) Hash {
	b64Sal := base64Encode(sal)
	b64Hash := base64Encode(hash)

	b64 := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		config.memory,
		config.iterations,
		config.parallelism,
		b64Sal,
		b64Hash,
	)

	return b64
}

// encriptarAES encripta a senha pelo algoritmo AES atraves do Galois/Counter Mode.
func encriptarAES(senhaPlana, chave []byte, nonceSize uint) []byte {
	cifraAES, err := aes.NewCipher(chave)
	if err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	gcm, err := cipher.NewGCMWithNonceSize(cifraAES, int(nonceSize))
	if err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	return gcm.Seal(nonce, nonce, senhaPlana, nil)
}

// nolint:gomnd
// GerenciadorSenhaPadrão retorna o gerenciador padrão de senhas.
func GerenciadorSenhaPadrão() *Senha {
	argon2Config := Argon2Config{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	return NovaSenha("meMudeMeMudeMeMudeMeMudeMeMudeMe", argon2Config, 16)
}
