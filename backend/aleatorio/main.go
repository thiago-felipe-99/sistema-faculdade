package aleatorio

import (
	"fmt"
	"math/rand"

	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

// Palavra criar uma Palavra aleatório de um tamnaho fixo.
func Palavra(tamanho int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũçABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ") //nolint:lll

	s := make([]rune, tamanho)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))] //nolint:gosec
	}

	return string(s)
}

// CPF cria um CPF aleatório.
func CPF() entidades.CPF {
	maxCPF := 99999999999

	return fmt.Sprintf("%011d", rand.Intn(maxCPF)) //nolint: GoSec
}
