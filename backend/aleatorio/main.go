package aleatorio

import (
	"fmt"
	"math/rand"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

// Palavra criar uma Palavra aleatório de um tamnaho fixo.
func Palavra(tamanho int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũçABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ") //nolint:lll

	rand.Seed(time.Now().UnixNano())

	s := make([]rune, tamanho)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))] //nolint: GoSec
	}

	return string(s)
}

//nolint: gomnd
// CPF cria um CPF aleatório.
func CPF() entidades.CPF {
	rand.Seed(time.Now().UnixNano())

	maxCPF := 999999999
	CPFSemDigitos := fmt.Sprintf("%09d", rand.Intn(maxCPF)) //nolint: GoSec

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
