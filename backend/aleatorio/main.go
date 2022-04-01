// Package aleatorio contém funções aleatórias necessárias para aplicação.
package aleatorio

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// ErroTamanhoInválido é um erro do pacote.
var ErroTamanhoInválido = erros.Padrão{
	Mensagem: "Tamanho da estrutura aleatória é inválida",
	Código:   "ALEATÓRIO-[1]",
}

// Número retorna um inteiro aleatório de [0,n).
func Número(n uint) uint {
	if n <= 0 {
		panic(ErroTamanhoInválido.Error())
	}

	número, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	return uint(número.Uint64())
}

// Palavra criar uma Palavra aleatório de um tamanho fixo.
func Palavra(tamanho uint) string {
	if tamanho <= 0 {
		panic(ErroTamanhoInválido.Error())
	}

	letters := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũçABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ")
	lettersLen := len(letters)

	s := make([]rune, tamanho)
	for i := range s {
		s[i] = letters[Número(uint(lettersLen))]
	}

	return string(s)
}

// CPF cria um CPF aleatório.
func CPF() string {
	const maxCPF = 999999999 + 1
	CPFSemDigitos := fmt.Sprintf("%09d", Número(maxCPF))

	digito1 := 0
	digito2 := 0

	for índice, número := range CPFSemDigitos {
		digito1 += int(número-'0') * (9 - índice%10)     //nolint: gomnd
		digito2 += int(número-'0') * (9 - (índice+1)%10) //nolint: gomnd
	}

	digito1 = (digito1 % 11) % 10 //nolint: gomnd
	digito2 += digito1 * 9        //nolint: gomnd
	digito2 = (digito2 % 11) % 10 //nolint: gomnd

	cpf := ""
	for índice := 8; índice >= 0; índice-- {
		cpf += string(CPFSemDigitos[índice])
	}

	cpf = fmt.Sprintf("%s%d%d", cpf, digito1, digito2)

	return cpf
}

// Bytes retorna uma slice de bytes aleatório de tamanho n.
func Bytes(n uint32) []byte {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		panic(erros.ErroExterno(err))
	}

	return bytes
}

// Senha retorna uma senha aleatória válida na aplicação.
func Senha() string {
	senha := ""

	letrasMinúsculas := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũç")
	letrasMaiúsculas := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ")
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
