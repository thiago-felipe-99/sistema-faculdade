package logs

import (
	"io"
	"log"
	"os"
)

type Log struct {
	Informação *log.Logger
	Aviso      *log.Logger
	Erro       *log.Logger
}

func NovoLog(out io.Writer) *Log {
	return &Log{
		Informação: log.New(out, "INFORMAÇÃO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Aviso:      log.New(out, "AVISO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Erro:       log.New(out, "ERRO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

type Arquivos struct {
	Pessoa         io.Writer
	Curso          io.Writer
	Aluno          io.Writer
	Professor      io.Writer
	Administrativo io.Writer
	Matéria        io.Writer
	Turma          io.Writer
}

//nolint:funlen
func AbrirArquivos(defaultDir string) *Arquivos {
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

	return &Arquivos{
		Pessoa:         pessoa,
		Curso:          curso,
		Aluno:          aluno,
		Professor:      professor,
		Administrativo: administrativo,
		Matéria:        matéria,
		Turma:          turma,
	}
}
