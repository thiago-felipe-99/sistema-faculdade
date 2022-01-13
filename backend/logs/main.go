package logs

import (
	"io"
	"log"
	"os"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

var ErroNívelInválido = &erros.Padrão{
	Mensagem: "O nível escolhido é inválido",
	Código:   "LOG-[1]",
}

const (
	NívelPanic uint = iota
	NívelErro
	NívelAviso
	NívelInfo
	NívelDebug
)

type Log struct {
	outPanic      *log.Logger
	outErro       *log.Logger
	outAviso      *log.Logger
	outInformação *log.Logger
	outDebug      *log.Logger
	Nível         uint
}

func (log *Log) Panic(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outPanic.Panicln(imprimir...)
}

func (log *Log) Erro(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outErro.Println(imprimir...)
}

func (log *Log) Aviso(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outAviso.Println(imprimir...)
}

func (log *Log) Informação(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outInformação.Println(imprimir...)
}

func (log *Log) Debug(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outDebug.Println(imprimir...)
}

func NovoLog(out io.Writer, nível uint) (*Log, *erros.Aplicação) {
	if nível < NívelPanic || nível > NívelDebug {
		return nil, erros.Novo(ErroNívelInválido, nil, nil)
	}

	return &Log{
		outPanic:      log.New(out, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile),
		outErro:       log.New(out, "ERRO: ", log.Ldate|log.Ltime|log.Lshortfile),
		outAviso:      log.New(out, "AVISO: ", log.Ldate|log.Ltime|log.Lshortfile),
		outInformação: log.New(out, "INFORMAÇÃO: ", log.Ldate|log.Ltime|log.Lshortfile),
		outDebug:      log.New(out, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		Nível:         nível,
	}, nil
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

func AbrirArquivos(defaultDir string) *Arquivos {
	const flags = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	const mode os.FileMode = 0666

	pessoa, err := os.OpenFile(defaultDir+"pessoaLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	curso, err := os.OpenFile(defaultDir+"cursoLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	aluno, err := os.OpenFile(defaultDir+"alunoLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	professor, err := os.OpenFile(defaultDir+"professorLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	administrativo, err := os.OpenFile(defaultDir+"administrativoLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	matéria, err := os.OpenFile(defaultDir+"matériaLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
	}

	turma, err := os.OpenFile(defaultDir+"turmaLogs.txt", flags, mode)
	if err != nil {
		log.Panic(err)
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

type Entidades struct {
	Pessoa         *Log
	Curso          *Log
	Aluno          *Log
	Professor      *Log
	Administrativo *Log
	Matéria        *Log
	Turma          *Log
}

func NovoLogEntidades(arquivos *Arquivos, nível uint) *Entidades {
	pessoa, err := NovoLog(arquivos.Pessoa, nível)
	if err != nil {
		log.Panicln(err)
	}

	curso, err := NovoLog(arquivos.Curso, nível)
	if err != nil {
		log.Panicln(err)
	}

	aluno, err := NovoLog(arquivos.Aluno, nível)
	if err != nil {
		log.Panicln(err)
	}

	professor, err := NovoLog(arquivos.Professor, nível)
	if err != nil {
		log.Panicln(err)
	}

	administrativo, err := NovoLog(arquivos.Administrativo, nível)
	if err != nil {
		log.Panicln(err)
	}

	matéria, err := NovoLog(arquivos.Matéria, nível)
	if err != nil {
		log.Panicln(err)
	}

	turma, err := NovoLog(arquivos.Turma, nível)
	if err != nil {
		log.Panicln(err)
	}

	return &Entidades{
		Pessoa:         pessoa,
		Curso:          curso,
		Aluno:          aluno,
		Professor:      professor,
		Administrativo: administrativo,
		Matéria:        matéria,
		Turma:          turma,
	}
}
