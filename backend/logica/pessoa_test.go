package logica_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"thiagofelipe.com.br/sistema-faculdade-backend/aleatorio"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
	//nolint:revive,stylecheck
	// . "thiagofelipe.com.br/sistema-faculdade-backend/logica"
)

func criarPessoaAleatória() (string, string, time.Time, string) {
	nome := aleatorio.Palavra(rand.Intn(tamanhoMáximoDaPalavra))
	cpf := aleatorio.CPF()
	data := time.Now()
	senha := aleatorio.Palavra(rand.Intn(tamanhoMáximoDaPalavra))

	return nome, cpf, data, senha
}

func removerPessoa(id entidades.ID, t *testing.T) {
	erro := logica.Pessoa.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar a pessoa teste: %v", erro.Traçado())
	}
}

func adicionarPessoa(
	nome string,
	cpf string,
	dataDeNascimento time.Time,
	senha string,
	t *testing.T,
) {
	pessoaCriada, erro := logica.Pessoa.Criar(nome, cpf, dataDeNascimento, senha)
	if erro != nil {
		t.Fatalf("Erro ao criar a pessoa: %s", erro.Traçado())
	}

	pessoaSalva, erro := logica.Pessoa.Pegar(pessoaCriada.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa: %s", erro.Traçado())
	}

	if !reflect.DeepEqual(pessoaCriada, pessoaSalva) {
		t.Fatalf(
			"Erro ao criar a pessoa, queria %v, chegou %v",
			pessoaCriada,
			pessoaSalva,
		)
	}

	t.Cleanup(func() {
		removerPessoa(pessoaCriada.ID, t)
	})
}

func TestCriarPessoa(t *testing.T) {
	nome, cpf, dataDeNascimento, senha := criarPessoaAleatória()

	adicionarPessoa(nome, cpf, dataDeNascimento, senha, t)
}
