package aleatorio

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

var erroTamanhoInválido = fmt.Errorf("tamanho inválido")

// Número retorna um inteiro aleatório de [0,n).
func Número(n uint) uint {
	if n <= 0 {
		log.Panicln(erroTamanhoInválido)
	}

	número, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	return uint(número.Uint64())
}

// Palavra criar uma Palavra aleatório de um tamnaho fixo.
func Palavra(tamanho uint) string {
	if tamanho <= 0 {
		log.Panicln(erroTamanhoInválido)
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũçABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ") //nolint:lll
	lettersLen := len(letters)

	if lettersLen <= 0 {
		log.Panicln(erroTamanhoInválido)
	}

	s := make([]rune, tamanho)
	for i := range s {
		s[i] = letters[Número(uint(lettersLen))]
	}

	return string(s)
}

// nolint:gomnd
// CPF cria um CPF aleatório.
func CPF() string {
	const maxCPF = 999999999 + 1
	CPFSemDigitos := fmt.Sprintf("%09d", Número(maxCPF))

	digito1 := 0
	digito2 := 0

	for índice, número := range CPFSemDigitos {
		digito1 += int(número-'0') * (9 - índice%10)
		digito2 += int(número-'0') * (9 - (índice+1)%10)
	}

	digito1 = (digito1 % 11) % 10
	digito2 += digito1 * 9
	digito2 = (digito2 % 11) % 10

	cpf := ""
	for índice := 8; índice >= 0; índice-- {
		cpf += string(CPFSemDigitos[índice])
	}

	cpf = fmt.Sprintf("%s%d%d", cpf, digito1, digito2)

	return cpf
}

// Bytes retorna uma slice de bytes aleatório de tamanho n.
func Bytes(n uint32) []byte {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		log.Panicln(erros.ErroExterno(err))
	}

	return b
}

// Senha retorna uma senha aleatória válida na aplicação.
func Senha() string {
	senha := ""

	letrasMinúsculas := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũç") //nolint:lll
	letrasMaiúsculas := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ") //nolint:lll
	tamanhoLetras := uint(len(letrasMaiúsculas))

	for índice := 0; índice < 4; índice++ {
		senha += string(letrasMaiúsculas[Número(tamanhoLetras)])
		senha += string(letrasMinúsculas[Número(tamanhoLetras)])
	}

	caractersEspeciais := []rune("@#$%^&-+=()")

	senha += fmt.Sprintf(
		"%s%d%c",
		senha,
		Número(tamanhoLetras),
		caractersEspeciais[Número(uint(len(caractersEspeciais)))],
	)

	return senha
}
