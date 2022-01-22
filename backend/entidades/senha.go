package entidades

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

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
func (senha *Senha) ÉIgual(
	senhaPlana string,
	hashOriginal Hash,
) (bool, *erros.Aplicação) {
	hashAES, err := base64Decodificar(hashOriginal)
	if err != nil {
		return false, erros.Novo(ErroVerificarSenhaHash, err, nil)
	}

	hashArgon2id, err := desencriptarAES(hashAES, senha.chave, senha.nonceSize)
	if err != nil {
		return false, erros.Novo(ErroVerificarSenhaHash, err, nil)
	}

	hashSHA512 := gerarHashSHA3_512([]byte(senhaPlana))

	return verificarSenhaHashArgon2id(string(hashSHA512), string(hashArgon2id))
}

// ÉVálida verifica se a senha cumpri os requisitos de uma senha forte.
func (senha *Senha) ÉVálida(senhaPlana string) bool {
	tamanho := utf8.RuneCountInString(senhaPlana)
	if tamanho < 8 || tamanho > 255 {
		return false
	}

	números := regexp.MustCompile(`[0-9]`)
	if !números.MatchString(senhaPlana) {
		return false
	}

	letrasMinúsculas := regexp.MustCompile(`\p{Ll}`)
	if !letrasMinúsculas.MatchString(senhaPlana) {
		return false
	}

	letrasMaiúsculas := regexp.MustCompile(`\p{Lu}`)
	if !letrasMaiúsculas.MatchString(senhaPlana) {
		return false
	}

	caractersEspeciasis := regexp.MustCompile(`[@#$%^&\-+=()]`)
	if !caractersEspeciasis.MatchString(senhaPlana) {
		return false
	}

	espaço := regexp.MustCompile(` `)

	return !espaço.MatchString(senhaPlana)
}

// GerarHash retorna a senhaPlana hasheada no algoritmo padrão.
func (senha *Senha) GerarHash(senhaPlana string) Hash {
	hashSHA512 := gerarHashSHA3_512([]byte(senhaPlana))

	hashArgon2id, saltArgon2id := gerarHashArgon2id(hashSHA512, senha.argon2)

	hashB64 := base64CodificarArgon2id(hashArgon2id, saltArgon2id, senha.argon2)

	hashAES := encriptarAES([]byte(hashB64), senha.chave, senha.nonceSize)

	hash := base64Codificar(hashAES)

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

	hash = gerarHashArgon2idComSal(senhaPlana, sal, config)

	return hash, sal
}

// gerarHashArgon2idComSal gera o hash de uma senha plana pelo algoritmo
// argon2id.
func gerarHashArgon2idComSal(senhaPlana, sal []byte, config Argon2Config) []byte {
	hash := argon2.IDKey(
		senhaPlana,
		sal,
		config.iterations,
		config.memory,
		config.parallelism,
		config.keyLength,
	)

	return hash
}

// verificarSenhaHashArgon2id verifica se a senha plana é igual o hash pelo
// algoritmo argon2id.
func verificarSenhaHashArgon2id(
	senhaPlana string,
	hashOriginal Hash,
) (bool, *erros.Aplicação) {
	hash, sal, config, erro := base64DecodificarArgon2id(hashOriginal)
	if erro != nil {
		return false, erros.Novo(ErroVerificarSenhaHash, erro, nil)
	}

	hashTeste := gerarHashArgon2idComSal([]byte(senhaPlana), sal, *config)

	if subtle.ConstantTimeCompare(hash, hashTeste) == 1 {
		return true, nil
	}

	return false, nil
}

// base64Encode retorna o slice de bytes codificado na base64.
func base64Codificar(bytes []byte) string {
	return base64.RawStdEncoding.EncodeToString(bytes)
}

// base64Decodificar retorna a string decodificada na base64.
func base64Decodificar(codificado string) ([]byte, *erros.Aplicação) {
	decodificado, err := base64.RawStdEncoding.Strict().DecodeString(codificado)
	if err != nil {
		return nil, erros.Novo(ErroDecodificarBase64, nil, err)
	}

	return decodificado, nil
}

// base64CodificarArgon2id retorna o hash de um algoritmo argon2id codificado na
// base64.
func base64CodificarArgon2id(senha, sal []byte, config Argon2Config) Hash {
	codificadoSal := base64Codificar(sal)
	codiciadoSenhaPlana := base64Codificar(senha)

	codificado := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		config.memory,
		config.iterations,
		config.parallelism,
		codificadoSal,
		codiciadoSenhaPlana,
	)

	return codificado
}

// base64DecodificarArgon2id retorna o hash, sal e as configurações usadas em um
// hash codificado na base 64 e gerado através do Argon2id.
func base64DecodificarArgon2id(hashB64 string) (
	senha, sal []byte,
	config *Argon2Config,
	erro *erros.Aplicação,
) {
	valores := strings.Split(hashB64, "$")
	if len(valores) != 6 { // nolint: gomnd
		return nil, nil, nil, erros.Novo(ErroDecodificarHashInválido, nil, nil)
	}

	if valores[1] != "argon2id" {
		return nil, nil, nil, erros.Novo(ErroDecodificarHashInválido, nil, nil)
	}

	var version int

	_, erroExterno := fmt.Sscanf(valores[2], "v=%d", &version)
	if erroExterno != nil {
		return nil, nil, nil, erros.Novo(ErroDecodificarArgon2id, nil, erroExterno)
	}

	if version != argon2.Version {
		return nil, nil, nil, erros.Novo(ErroDecodificarArgon2idVersão, nil, nil)
	}

	var (
		memory, iterations uint32
		parallelism        uint8
	)

	_, erroExterno = fmt.Sscanf(
		valores[3],
		"m=%d,t=%d,p=%d",
		&memory,
		&iterations,
		&parallelism,
	)
	if erroExterno != nil {
		return nil, nil, nil, erros.Novo(ErroDecodificarArgon2id, nil, erroExterno)
	}

	sal, erro = base64Decodificar(valores[4])
	if erro != nil {
		return nil, nil, nil, erros.Novo(ErroDecodificarArgon2id, erro, nil)
	}

	senha, erro = base64Decodificar(valores[5])
	if erro != nil {
		return nil, nil, nil, erros.Novo(ErroDecodificarArgon2id, erro, nil)
	}

	config = &Argon2Config{
		memory:      memory,
		iterations:  iterations,
		parallelism: parallelism,
		saltLength:  uint32(len(sal)),
		keyLength:   uint32(len(senha)),
	}

	return senha, sal, config, nil
}

// encriptarAES encripta a senha pelo algoritmo AES atraves do Galois/Counter
// Mode.
func encriptarAES(senhaPlana, chave []byte, nonceSize uint) []byte {
	cifraAES, err := aes.NewCipher(chave)
	if err != nil {
		erro := erros.ErroExterno(err)
		log.Println(erro)
		panic(erro)
	}

	gcm, err := cipher.NewGCMWithNonceSize(cifraAES, int(nonceSize))
	if err != nil {
		erro := erros.ErroExterno(err)
		log.Println(erro)
		panic(erro)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		erro := erros.ErroExterno(err)
		log.Println(erro)
		panic(erro)
	}

	return gcm.Seal(nonce, nonce, senhaPlana, nil)
}

// desencriptarAES desencripta a senha pelo algoritmo AES atraves do
// Galois/Counter Mode.
func desencriptarAES(senhaCifrada, chave []byte, nonceSize uint) (
	[]byte,
	*erros.Aplicação,
) {
	cifraAES, err := aes.NewCipher(chave)
	if err != nil {
		return nil, erros.Novo(ErroDesencriptarAES, nil, err)
	}

	gcm, err := cipher.NewGCMWithNonceSize(cifraAES, int(nonceSize))
	if err != nil {
		return nil, erros.Novo(ErroDesencriptarAES, nil, err)
	}

	if len(senhaCifrada) < gcm.NonceSize() {
		return nil, erros.Novo(ErroDesencriptarAESNonceSize, nil, nil)
	}

	nonce, senhaCifrada := senhaCifrada[:nonceSize], senhaCifrada[nonceSize:]

	senhaPlana, err := gcm.Open(nil, nonce, senhaCifrada, nil)
	if err != nil {
		return nil, erros.Novo(ErroDesencriptarAES, nil, err)
	}

	return senhaPlana, nil
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
