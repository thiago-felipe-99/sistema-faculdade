package data

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/env"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

var ambiente = env.PegandoVariáveisDeAmbiente()

func criarConexãoMariaDB() *sql.DB {
	//nolint:exhaustivestruct
	config := mysql.Config{
		User:                 "Teste",
		Passwd:               "Teste",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:" + ambiente.Portas.BDAdministrativo,
		DBName:               "Teste",
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
	}

	conexão, erro := mariadb.NovoBD(config.FormatDSN())
	if erro != nil {
		log.Fatalf("Erro ao configurar o banco de dados: %s", erro)
	}

	erroPing := conexão.Ping()
	if erroPing != nil {
		log.Fatalf("Erro ao conectar o banco de dados: %s", erroPing)
	}

	return conexão
}

func TestDataPadrão(t *testing.T) {
	bd := criarConexãoMariaDB()
	logFiles := logs.AbrirArquivos("./logs/")
	log := logs.NovoLogEntidades(logFiles, logs.NívelDebug)

	data := DataPadrão(log, bd)

	tipos := map[string]struct{ quer, recebou string }{
		"Pessoa":         {"mariadb.PessoaBD", fmt.Sprintf("%T", data.Pessoa)},
		"Curso":          {"mariadb.CursoBD", fmt.Sprintf("%T", data.Curso)},
		"Aluno":          {"mariadb.AlunoBD", fmt.Sprintf("%T", data.Aluno)},
		"Professor":      {"mariadb.ProfessorBD", fmt.Sprintf("%T", data.Professor)},
		"Administrativo": {"mariadb.AdministrativoBD", fmt.Sprintf("%T", data.Administrativo)},
		"Matéria":        {"mongodb.MatériaBD", fmt.Sprintf("%T", data.Matéria)},
		"Turma":          {"mongodb.TurmaBD", fmt.Sprintf("%T", data.Turma)},
	}

	for chave, valor := range tipos {
		if valor.quer != valor.recebou {
			t.Fatalf(
				"A data da entidade %s quer: %s, porém recebeu: %s",
				chave, valor.quer, valor.recebou,
			)
		}
	}
}
