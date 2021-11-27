package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mariadb"
)

func newData() *data.Data {
	pessoaFile, err := os.OpenFile(
		"pessoaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	cursoFile, err := os.OpenFile(
		"cursoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	alunoFile, err := os.OpenFile(
		"alunoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	professorFile, err := os.OpenFile(
		"professorLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	administrativoFile, err := os.OpenFile(
		"administrativoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	MariaDBPessoa := mariadb.Pessoa{
		Connection: *mariadb.NewConnection(pessoaFile),
	}

	MariaDBCurso := mariadb.Curso{
		Connection: *mariadb.NewConnection(cursoFile),
	}

	MariaDBAluno := mariadb.Aluno{
		Connection: *mariadb.NewConnection(alunoFile),
	}

	MariaDBProfessor := mariadb.Professor{
		Connection: *mariadb.NewConnection(professorFile),
	}

	MariaDBAdministrativo := mariadb.Administrativo{
		Connection: *mariadb.NewConnection(administrativoFile),
	}

	return &data.Data{
		Pessoa:         MariaDBPessoa,
		Curso:          MariaDBCurso,
		Aluno:          MariaDBAluno,
		Professor:      MariaDBProfessor,
		Administrativo: MariaDBAdministrativo,
	}
}

func main() {
	r := gin.Default()

	Data := newData()

	r.GET("/ping", func(c *gin.Context) {
		Data.Pessoa.Get(uuid.New())
		Data.Curso.Get(uuid.New())
		Data.Aluno.Get(uuid.New())
		Data.Professor.Get(uuid.New())
		Data.Administrativo.Get(uuid.New())
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
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
