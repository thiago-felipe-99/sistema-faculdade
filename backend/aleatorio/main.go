package aleatorio

import "math/rand"

// Palavra criar uma Palavra aleatório de um tamnaho fixo.
func Palavra(tamanho int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzáéíóúâêîôûãẽĩõũçABCDEFGHIJKLMNOPQRSTUVWXYZÁÉÍÓÚÂÊÎÔÛÃẼĨÕŨÇ") //nolint:lll

	s := make([]rune, tamanho)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))] //nolint:gosec
	}

	return string(s)
}
