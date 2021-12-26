package mariadb

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	. "thiagofelipe.com.br/sistema-faculdade-backend/data/erros"
	"thiagofelipe.com.br/sistema-faculdade-backend/entidades"
)

func criarPessoaAleatória() *entidades.Pessoa {
	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	var pessoa = &entidades.Pessoa{
		ID:               entidades.NovoID(),
		Nome:             "Teste Certo",
		CPF:              fmt.Sprintf("%011d", rand.Intn(99999999999)),
		DataDeNascimento: dataAgora,
		Senha:            "Senha",
	}

	return pessoa
}

func adicionarPessoa(pessoa *entidades.Pessoa, t *testing.T) {

	erro := pessoaBD.Inserir(pessoa)
	if erro != nil {
		t.Fatalf("Erro ao inserir a pessoa no banco de dados: %s", erro.Error())
	}

	pessoaSalva, erro := pessoaBD.Pegar(pessoa.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(pessoa, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoa,
			pessoaSalva,
		)
	}

	t.Cleanup(func() {
		removerPessoa(pessoa.ID, t)
	})
}

func removerPessoa(id entidades.ID, t *testing.T) {
	erro := pessoaBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", erro.Error())
	}
}

func TestInserirPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)
}

func TestInserirPessoa_duplicadoID(t *testing.T) {
	texto := `Duplicate entry.*PRIMARY`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	pessoaTeste.CPF = fmt.Sprintf("%011d", rand.Intn(99999999999))

	erro := pessoaBD.Inserir(pessoaTeste)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestInserirPessoa_duplicadoCPF(t *testing.T) {
	texto := `Duplicate entry.*CPF`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	defer removerPessoa(pessoaTeste.ID, t)

	pessoaTeste.ID = entidades.NovoID()

	erro := pessoaBD.Inserir(pessoaTeste)

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestAtualizarPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	pessoaTeste.Nome = "Novo Nome"
	pessoaTeste.CPF = "00000000000"
	pessoaTeste.DataDeNascimento = dataAgora.AddDate(0, 0, 30)
	pessoaTeste.Senha = "Senha nova"

	erro := pessoaBD.Atualizar(pessoaTeste.ID, pessoaTeste)
	if erro != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", erro.Error())
	}

	pessoaSalva, erro := pessoaBD.Pegar(pessoaTeste.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoaTeste,
			pessoaSalva,
		)
	}
}

func TestAtualizarPessoa_duplicadoID(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	erro := pessoaBD.Atualizar(entidades.NovoID(), pessoaTeste)
	if erro != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", erro.Error())
	}
}

func TestAtualizarPessoa_duplicadoCPF(t *testing.T) {
	texto := `Duplicate entry.*CPF`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste1 := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste1, t)

	pessoaTeste2 := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste2, t)

	erro := pessoaBD.Atualizar(pessoaTeste1.ID, pessoaTeste2)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: %s, chegou %s",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestPegarPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	pessoaSalva, erro := pessoaBD.Pegar(pessoaTeste.ID)
	if erro != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", erro.Error())
	}

	if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoaTeste,
			pessoaSalva,
		)
	}
}

func TestPegarPessoa_inválidoID(t *testing.T) {
	_, erro := pessoaBD.Pegar(entidades.NovoID())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
		t.Fatalf(
			"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
			ErroPessoaNãoEncontrada.Error(),
			erro.Error(),
		)
	}
}

func TestPegarPessoa_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	_, erro := pessoaBDInválido.Pegar(entidades.NovoID())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}

func TestDeletarPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adicionarPessoa(pessoaTeste, t)

	removerPessoa(pessoaTeste.ID, t)

	_, erro := pessoaBD.Pegar(pessoaTeste.ID)
	if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
		t.Fatalf(
			"Deveria retonar um erro de pessoa não encontrada, retonou %s",
			erro,
		)
	}
}

func TestDeletarPessoa_invalídoID(t *testing.T) {
	id := entidades.NovoID()

	removerPessoa(id, t)

	_, erro := pessoaBD.Pegar(id)
	if erro == nil || !erro.ÉPadrão(ErroPessoaNãoEncontrada) {
		t.Fatalf(
			"Deveria retonar um erro de pessoa não encontrada, retonou %s",
			erro,
		)
	}
}

func TestDeletarPessoa_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := pessoaBDInválido.Deletar(entidades.NovoID())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
			texto,
			erro.ErroExterno.Error(),
		)
	}
}
