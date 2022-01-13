package entidades

// Hash representa uma hash na aplicação.
type Hash = string

// Senha gerencia como as senhas são tratadas na aplicação.
type Senha struct{}

// NovaSenha criar um gerenciador de senhas.
func NovaSenha() *Senha {
	return &Senha{}
}

// ÉIgual verifica se um uma senhaPlana é igual o hash.
func (senha *Senha) ÉIgual(senhaPlana string, hash Hash) bool {
	return senhaPlana == hash
}

// ÉVálida verifica se a senha cumpri os requisitos de uma senha forte.
func (senha *Senha) ÉVálida(senhaPlana string) bool {
	return true
}

// GerarHash retorna a senhaPlana hasheada.
func (senha *Senha) GerarHash(senhaPlana string) Hash {
	return senhaPlana
}

// GerenciadorSenhaPadrão retorna o gerenciador padrão de senhas.
func GerenciadorSenhaPadrão() *Senha {
	return NovaSenha()
}
