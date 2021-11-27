package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"thiagofelipe.com.br/sistema-faculdade/data"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mariadb"
	"thiagofelipe.com.br/sistema-faculdade/data/database/mongodb"
)

type logFiles struct {
	pessoa         io.Writer
	curso          io.Writer
	aluno          io.Writer
	professor      io.Writer
	administrativo io.Writer
	matéria        io.Writer
	turma          io.Writer
}

func openLogFiles() *logFiles {

	defaultDir := "./logs/data/"

	pessoa, err := os.OpenFile(
		defaultDir+"pessoaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	curso, err := os.OpenFile(
		defaultDir+"cursoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	aluno, err := os.OpenFile(
		defaultDir+"alunoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	professor, err := os.OpenFile(
		defaultDir+"professorLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	administrativo, err := os.OpenFile(
		defaultDir+"administrativoLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	matéria, err := os.OpenFile(
		defaultDir+"matériaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	turma, err := os.OpenFile(
		defaultDir+"turmaLogs.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}

	return &logFiles{
		pessoa:         pessoa,
		curso:          curso,
		aluno:          aluno,
		professor:      professor,
		administrativo: administrativo,
		matéria:        matéria,
		turma:          turma,
	}
}

func newData() *data.Data {
	logFiles := openLogFiles()

	MariaDBPessoa := mariadb.Pessoa{
		Connection: *mariadb.NewConnection(logFiles.pessoa),
	}

	MariaDBCurso := mariadb.Curso{
		Connection: *mariadb.NewConnection(logFiles.curso),
	}

	MariaDBAluno := mariadb.Aluno{
		Connection: *mariadb.NewConnection(logFiles.aluno),
	}

	MariaDBProfessor := mariadb.Professor{
		Connection: *mariadb.NewConnection(logFiles.professor),
	}

	MariaDBAdministrativo := mariadb.Administrativo{
		Connection: *mariadb.NewConnection(logFiles.administrativo),
	}

	MongoDBMatéria := mongodb.Matéria{
		Connection: *mongodb.NewConnection(logFiles.matéria),
	}

	MongoDBTurma := mongodb.Turma{
		Connection: *mongodb.NewConnection(logFiles.turma),
	}

	return &data.Data{
		Pessoa:         MariaDBPessoa,
		Curso:          MariaDBCurso,
		Aluno:          MariaDBAluno,
		Professor:      MariaDBProfessor,
		Administrativo: MariaDBAdministrativo,
		Matéria:        MongoDBMatéria,
		Turma:          MongoDBTurma,
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
