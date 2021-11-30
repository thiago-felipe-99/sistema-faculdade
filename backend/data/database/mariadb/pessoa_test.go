package mariadb

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/env"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

var bd *PessoaBD

var bdInválido *PessoaBD

func TestMain(m *testing.M) {
	ambiente := env.PegandoVariáveisDeAmbiente()

	config := mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	connexão, erro := NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	erroPing := connexão.Ping()
	if erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	bd = &PessoaBD{
		Conexão:      *NovaConexão(os.Stderr, connexão),
		NomeDaTabela: "Pessoa",
	}

	bdInválido = &PessoaBD{
		Conexão:      *NovaConexão(os.Stderr, connexão),
		NomeDaTabela: "PessoaErrada",
	}

	código := m.Run()

	query := "DELETE FROM Pessoa"
	bd.BD.Exec(query)

	os.Exit(código)
}

func criarPessoaAleatória() *data.Pessoa {
	rand.Seed(time.Now().UnixNano())

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	var pessoa = &data.Pessoa{
		ID:               uuid.New(),
		Nome:             "Teste Certo",
		CPF:              fmt.Sprintf("%011d", rand.Intn(99999999999)),
		DataDeNascimento: dataAgora,
		Senha:            "Senha",
	}

	return pessoa
}

func adiconarPessoa(pessoa *data.Pessoa, t *testing.T) {

	erro := bd.Inserir(pessoa)
	if erro != nil {
		t.Fatalf("Erro ao inserir a pessoa no banco de dados: %s", erro.Error())
	}

	pessoaSalva, erro := bd.Pegar(pessoa.ID)
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
	erro := bd.Deletar(id)
	if erro != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", erro.Error())
	}
}

func TestInserir(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)
}

func TestInserir_duplicadoID(t *testing.T) {
	padrão, erroRegex := regexp.Compile(`Duplicate entry.*PRIMARY`)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	pessoaTeste.CPF = fmt.Sprintf("%011d", rand.Intn(99999999999))

	erro := bd.Inserir(pessoaTeste)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'PRIMARY', chegou %s",
			erro.ErroExterno.Error(),
		)
	}
}

func TestInserir_duplicadoCPF(t *testing.T) {
	padrão, erroRegex := regexp.Compile(`Duplicate entry.*CPF`)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	defer removerPessoaTeste(pessoaTeste.ID, t)

	pessoaTeste.ID = uuid.New()

	erro := bd.Inserir(pessoaTeste)

	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'CPF', chegou %s",
			erro.ErroExterno.Error(),
		)
	}
}

func TestAtualizar(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	pessoaTeste.Nome = "Novo Nome"
	pessoaTeste.CPF = "00000000000"
	pessoaTeste.DataDeNascimento = dataAgora.AddDate(0, 0, 30)
	pessoaTeste.Senha = "Senha nova"

	erro := bd.Atualizar(pessoaTeste.ID, pessoaTeste)
	if erro != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", erro.Error())
	}

	pessoaSalva, erro := bd.Pegar(pessoaTeste.ID)
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

func TestAtualizar_duplicadoID(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	erro := bd.Atualizar(uuid.New(), pessoaTeste)
	if erro != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", erro.Error())
	}
}

func TestAtualizar_duplicadoCPF(t *testing.T) {
	padrão, erroRegex := regexp.Compile(`Duplicate entry.*CPF`)
	if erroRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste1 := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste1, t)

	pessoaTeste2 := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste2, t)

	erro := bd.Atualizar(pessoaTeste1.ID, pessoaTeste2)
	if erro == nil || erro.ErroExterno == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !padrão.MatchString(erro.ErroExterno.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'CPF', chegou %s",
			erro.ErroExterno.Error(),
		)
	}
}

func TestPegar(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	pessoaSalva, erro := bd.Pegar(pessoaTeste.ID)
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

func TestPegar_inválidoID(t *testing.T) {
	_, erro := bd.Pegar(uuid.New())

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

func TestPegar_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	_, erro := bdInválido.Pegar(uuid.New())

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

func TestDeletar(t *testing.T) {
	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, t)

	removerPessoaTeste(pessoaTeste.ID, t)
}

func TestDeletar_invalídoID(t *testing.T) {
	removerPessoaTeste(uuid.New(), t)
}

func TestDeletar_tabelaInválida(t *testing.T) {
	texto := `Table .* doesn't exist`
	padrão, erroRegex := regexp.Compile(texto)
	if erroRegex != nil {
		t.Fatalf("Erro ao compilar o regex: %s", texto)
	}

	erro := bdInválido.Deletar(uuid.New())

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
