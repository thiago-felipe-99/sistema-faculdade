package mariadb

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/errors"
)

var config = mysql.Config{
	User:                 "Teste",
	Passwd:               "Teste",
	Net:                  "tcp",
	Addr:                 "localhost:9000",
	DBName:               "Teste",
	AllowNativePasswords: true,
	ParseTime:            true,
}

func connection(t *testing.T) *PessoaDB {

	db, err := NewDB(config.FormatDSN())
	if err != nil {
		t.Fatalf("Erro ao configurar o banco de dados: %s", err)
	}

	t.Cleanup(func() {
		query := "DELETE FROM Pessoa"
		db.Exec(query)
	})

	return &PessoaDB{
		Connection: *NewConnection(os.Stderr, db),
		TableName:  "Pessoa",
	}
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

func adiconarPessoa(pessoa *data.Pessoa, db *PessoaDB, t *testing.T) {

	err := db.Insert(pessoa)
	if err != nil {
		t.Fatalf("Erro ao inserir a pessoa no banco de dados: %s", err.Error())
	}

	pessoaSalva, err := db.Get(pessoa.ID)
	if err != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", err.Error())
	}

	if !reflect.DeepEqual(pessoa, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoa,
			pessoaSalva,
		)
	}

	t.Cleanup(func() {
		err := db.Delete(pessoa.ID)
		if err != nil {
			t.Fatalf("Erro ao tentar deletar o usuário teste: %v", err.Error())
		}
	})
}

func removerPessoaTeste(id id, db *PessoaDB, t *testing.T) {
	err := db.Delete(id)
	if err != nil {
		t.Fatalf("Erro ao tentar deletar o usuário teste: %v", err.Error())
	}
}

func TestInsert(t *testing.T) {
	db := connection(t)

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)
}

func TestInsert_duplicateID(t *testing.T) {
	db := connection(t)

	pattern, errRegex := regexp.Compile(`Duplicate entry.*PRIMARY`)
	if errRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	pessoaTeste.CPF = fmt.Sprintf("%011d", rand.Intn(99999999999))

	err := db.Insert(pessoaTeste)
	if err == nil || err.SystemError == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !pattern.MatchString(err.SystemError.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'PRIMARY', chegou %s",
			err.SystemError.Error(),
		)
	}
}

func TestInsert_duplicateCPF(t *testing.T) {
	db := connection(t)

	pattern, errRegex := regexp.Compile(`Duplicate entry.*CPF`)
	if errRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	pessoaTeste.ID = uuid.New()

	err := db.Insert(pessoaTeste)

	if err == nil || err.SystemError == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !pattern.MatchString(err.SystemError.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'CPF', chegou %s",
			err.SystemError.Error(),
		)
	}
}

func TestUpdate(t *testing.T) {
	db := connection(t)

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	dataAgora := time.Now().UTC()
	dataAgora = dataAgora.Truncate(24 * time.Hour)

	pessoaTeste.Nome = "Novo Nome"
	pessoaTeste.CPF = "00000000000"
	pessoaTeste.DataDeNascimento = dataAgora.AddDate(0, 0, 30)
	pessoaTeste.Senha = "Senha nova"

	err := db.Update(pessoaTeste.ID, pessoaTeste)
	if err != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", err.Error())
	}

	pessoaSalva, err := db.Get(pessoaTeste.ID)
	if err != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", err.Error())
	}

	if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoaTeste,
			pessoaSalva,
		)
	}
}

func TestUpdate_invalidID(t *testing.T) {
	db := connection(t)

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	pessoaTeste.ID = uuid.New()

	err := db.Update(pessoaTeste.ID, pessoaTeste)
	if err != nil {
		t.Fatalf("Erro ao atualizar a pessoa teste: %s", err.Error())
	}
}

func TestUpdate_duplicateCPF(t *testing.T) {
	db := connection(t)

	pattern, errRegex := regexp.Compile(`Duplicate entry.*CPF`)
	if errRegex != nil {
		t.Fatal("Erro ao compilar o regex")
	}

	pessoaTeste1 := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste1, db, t)

	pessoaTeste2 := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste2, db, t)

	err := db.Update(pessoaTeste1.ID, pessoaTeste2)
	if err == nil || err.SystemError == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !pattern.MatchString(err.SystemError.Error()) {
		t.Fatalf(
			"Erro ao inserir a pessoa queria: Duplicate entry for key 'CPF', chegou %s",
			err.SystemError.Error(),
		)
	}
}

func TestGet(t *testing.T) {
	db := connection(t)

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	pessoaSalva, err := db.Get(pessoaTeste.ID)
	if err != nil {
		t.Fatalf("Erro ao pegar a pessoa no banco de dados: %s", err.Error())
	}

	if !reflect.DeepEqual(pessoaTeste, pessoaSalva) {
		t.Fatalf(
			"Erro ao salvar a pessoa no banco de dados, queria %v, chegou %v",
			pessoaTeste,
			pessoaSalva,
		)
	}
}

func TestGet_invalidID(t *testing.T) {
	db := connection(t)

	_, err := db.Get(uuid.New())

	if err == nil || err.SystemError == nil {
		t.Fatalf("Não foi enviado erro do sistema")
	}

	if !err.IsDefault(errors.PessoaNãoEncontrada) {
		t.Fatalf(
			"Erro ao pegar pessoa no banco de dados, queria %v, chegou %v",
			errors.PessoaNãoEncontrada.Error(),
			err.Error(),
		)
	}
}

func TestDelete(t *testing.T) {
	db := connection(t)

	pessoaTeste := criarPessoaAleatória()

	adiconarPessoa(pessoaTeste, db, t)

	removerPessoaTeste(pessoaTeste.ID, db, t)
}

func TestDelete_invalidID(t *testing.T) {
	db := connection(t)

	removerPessoaTeste(uuid.New(), db, t)
}
