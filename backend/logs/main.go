package logs

import (
	"io"
	"log"
	"os"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

var ErroNívelInválido = &erros.Padrão{
	Mensagem: "O nível escolhido é inválido",
	Código:   "LOGS-[1]",
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
	log.outPanic.Println(imprimir...)
	panic(imprimir)
}

func (log *Log) Erro(imprimir ...interface{}) {
	if log.Nível < NívelErro {
		return
	}

	log.outErro.Println(imprimir...)
}

func (log *Log) Aviso(imprimir ...interface{}) {
	if log.Nível < NívelAviso {
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
	if log.Nível < NívelDebug {
		return
	}

	log.outDebug.Println(imprimir...)
}

func NovoLog(out io.Writer, nível uint) *Log {
	if nível < NívelPanic || nível > NívelDebug {
		panic(erros.Novo(ErroNívelInválido, nil, nil))
	}

	return &Log{
		outPanic:      log.New(out, "PANIC - ", log.Ldate|log.Ltime|log.Lshortfile),
		outErro:       log.New(out, "ERRO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outAviso:      log.New(out, "AVISO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outInformação: log.New(out, "INFORMAÇÃO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outDebug:      log.New(out, "DEBUG - ", log.Ldate|log.Ltime|log.Lshortfile),
		Nível:         nível,
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

func AbrirArquivos(defaultDir string) *Arquivos {
	const flags = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	const mode os.FileMode = 0o666

	const extension = ".log"

	pessoa, err := os.OpenFile(defaultDir+"Pessoa"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	curso, err := os.OpenFile(defaultDir+"Curso"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	aluno, err := os.OpenFile(defaultDir+"Aluno"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	professor, err := os.OpenFile(defaultDir+"Professor"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	administrativo, err := os.OpenFile(defaultDir+"Administrativo"+extension, flags, mode) //nolint:lll
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	matéria, err := os.OpenFile(defaultDir+"Matéria"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	turma, err := os.OpenFile(defaultDir+"Turma"+extension, flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
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
	pessoa := NovoLog(arquivos.Pessoa, nível)

	curso := NovoLog(arquivos.Curso, nível)

	aluno := NovoLog(arquivos.Aluno, nível)

	professor := NovoLog(arquivos.Professor, nível)

	administrativo := NovoLog(arquivos.Administrativo, nível)

	matéria := NovoLog(arquivos.Matéria, nível)

	turma := NovoLog(arquivos.Turma, nível)

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
