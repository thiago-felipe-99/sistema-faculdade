package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"thiagofelipe.com.br/sistema-faculdade-backend/data"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/database/mariadb"
	"thiagofelipe.com.br/sistema-faculdade-backend/data/database/mongodb"
	"thiagofelipe.com.br/sistema-faculdade-backend/logica"
	"thiagofelipe.com.br/sistema-faculdade-backend/logs"
)

//nolint:funlen
func newData() *data.Data {
	logFiles := logs.AbrirArquivos("./logs/data/")

	//nolint:exhaustivestruct
	config := mysql.Config{
		User:                 "Administrativo",
		Passwd:               "Administrativo",
		Net:                  "tcp",
		Addr:                 "localhost:9000",
		DBName:               "Administrativo",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dns := config.FormatDSN()
	log.Println(dns)

	bd, err := mariadb.NovoBD(config.FormatDSN())
	if err != nil {
		log.Panicln(err.Mensagem)
	}

	MariaDBPessoa := mariadb.PessoaBD{
		Conexão:      *mariadb.NovaConexão(logFiles.Pessoa, bd),
		NomeDaTabela: "Pessoa",
	}

	MariaDBCurso := mariadb.CursoBD{
		Conexão:                *mariadb.NovaConexão(logFiles.Curso, bd),
		NomeDaTabela:           "Curso",
		NomeDaTabelaSecundária: "CursoMatérias",
	}

	MariaDBAluno := mariadb.AlunoBD{
		Conexão:                *mariadb.NovaConexão(logFiles.Aluno, bd),
		NomeDaTabela:           "Aluno",
		NomeDaTabelaSecundária: "AlunoTurma",
	}

	MariaDBProfessor := mariadb.ProfessorBD{
		Conexão: *mariadb.NovaConexão(logFiles.Professor, bd),
	}

	MariaDBAdministrativo := mariadb.AdministrativoBD{
		Conexão: *mariadb.NovaConexão(logFiles.Administrativo, bd),
	}

	MariaDBMatéria := mongodb.MatériaBD{
		Connexão: *mongodb.NovaConexão(logFiles.Matéria),
	}

	MariaDBTurma := mongodb.TurmaBD{
		Connexão: *mongodb.NovaConexão(logFiles.Turma),
	}

	return &data.Data{
		Pessoa:         MariaDBPessoa,
		Curso:          MariaDBCurso,
		Aluno:          MariaDBAluno,
		Professor:      MariaDBProfessor,
		Administrativo: MariaDBAdministrativo,
		Matéria:        MariaDBMatéria,
		Turma:          MariaDBTurma,
	}
}

func prettyStruct(s ...interface{}) string {
	sJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Panicln(err.Error())
	}

	return string(sJSON)
}

func main() {
	r := gin.Default()

	data := newData()
	lógica := logica.NovaLógica(*data)

	r.GET("/ping", func(c *gin.Context) {
		pessoa, erro := lógica.Pessoa.Criar(
			"Thiago Felipe",
			"00000000000",
			time.Date(1999, 12, 8, 0, 0, 0, 0, time.UTC),
			"senha",
		)
		if erro != nil {
			log.Println(erro.Traçado())

			return
		}

		log.Println(prettyStruct(pessoa))

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"pessoa":  pessoa,
		})
	})

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
