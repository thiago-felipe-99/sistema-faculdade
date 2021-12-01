package mariadb

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/entidades"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

func criarPessoaAleatória() *entidades.Pessoa {
	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	var pessoa = &entidades.Pessoa{
		ID:               uuid.New(),
		Nome:             "Teste Certo",
		CPF:              fmt.Sprintf("%011d", rand.Intn(99999999999)),
		DataDeNascimento: dataAgora,
		Senha:            "Senha",
	}

	return pessoa
}

func adiconarPessoa(pessoa *entidades.Pessoa, t *testing.T) {

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
		removerPessoaTeste(pessoa.ID, t)
	})
}

func removerPessoaTeste(id id, t *testing.T) {
	erro := pessoaBD.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", erro.Error())
	}
}

func TestInserirPessoa(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)
}

func TestInserirPessoa_duplicadoID(t *testing.T) {
	texto := `Duplicate entry.*PRIMARY`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

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

	adiconarPessoa(pessoaTeste, t)

	defer removerPessoaTeste(pessoaTeste.ID, t)

	pessoaTeste.ID = uuid.New()

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

	adiconarPessoa(pessoaTeste, t)

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

	adiconarPessoa(pessoaTeste, t)

	erro := pessoaBD.Atualizar(uuid.New(), pessoaTeste)
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

	adiconarPessoa(pessoaTeste1, t)

	pessoaTeste2 := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste2, t)

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

	adiconarPessoa(pessoaTeste, t)

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
	_, erro := pessoaBD.Pegar(uuid.New())

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !erro.IsDefault(errors.PessoaNãoEncontrada) {
		t.Fatalf(
			"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
			errors.PessoaNãoEncontrada.Error(),
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

	_, erro := pessoaBDInválido.Pegar(uuid.New())

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

	adiconarPessoa(pessoaTeste, t)

	removerPessoaTeste(pessoaTeste.ID, t)
}

func TestDeletarPessoa_invalídoID(t *testing.T) {
	removerPessoaTeste(uuid.New(), t)
}

func TestDeletarPessoa_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := pessoaBDInválido.Deletar(uuid.New())

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
