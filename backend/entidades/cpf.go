package entidades

import "regexp"

// CPF representa o documento CPF(Cadatro De Pessoa Física) do Brasil.
type CPF = string

// parseCPF verifica se a string é um cpf válido por regex.
func parseCPF(cpf string) (CPF, bool) {
	cpfVálido := "00000000000"

	cpfRegra := regexp.
		MustCompile(`^([0-9]{3}\.[0-9]{3}\.[0-9]{3}-[0-9]{2}|[0-9]{11})$`)

	cpfs := cpfRegra.FindAllString(cpf, -1)
	if cpfs == nil || len(cpfs) != 1 {
		return cpfVálido, false
	}

	cpfPontos := regexp.MustCompile(`\.|-`)
	cpfVálido = cpfPontos.ReplaceAllString(cpfs[0], "")

	return cpfVálido, true
}

// nolint:gomnd
// verificarDígitoCPF verfica se os dígitos de verificação do CPF são válidos.
func verificarDígitoCPF(cpf string) bool {
	cpf, parse := parseCPF(cpf)
	if !parse {
		return parse
	}

	const CPFSemSigitos = 9

	cpfNúmeros := [CPFSemSigitos]int{}
	for índice := 0; índice < CPFSemSigitos; índice++ {
		cpfNúmeros[CPFSemSigitos-índice-1] = int(cpf[índice] - '0')
	}

	digito1 := 0
	digito2 := 0

	for índice, número := range cpfNúmeros {
		digito1 += número * (9 - índice%10)
		digito2 += número * (9 - (índice+1)%10)
	}

	digito1 = (digito1 % 11) % 10
	digito2 += digito1 * 9
	digito2 = (digito2 % 11) % 10

	digito1CPF := int(cpf[CPFSemSigitos] - '0')
	digito2CPF := int(cpf[CPFSemSigitos+1] - '0')

	return digito1 == digito1CPF && digito2 == digito2CPF
}

// ValidarCPF verifica se a string é um CPF.
func ValidarCPF(cpf string) (CPF, bool) {
	cpfVálido := "00000000000"

	cpf, parse := parseCPF(cpf)
	if !parse {
		return cpfVálido, false
	}

	dígitosVálidos := verificarDígitoCPF(cpf)
	if !dígitosVálidos {
		return cpfVálido, false
	}

	return cpf, true
}
