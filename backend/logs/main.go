// Package logs define como vai ser tratado os logs da aplicação.
package logs

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"thiagofelipe.com.br/sistema-faculdade-backend/erros"
)

// ErroNívelInválido da aplicação.
var ErroNívelInválido = &erros.Padrão{
	Mensagem: "O nível escolhido é inválido",
	Código:   "LOGS-[1]",
}

// Possíveis níveis de log.
const (
	NívelPanic uint = iota
	NívelErro
	NívelAviso
	NívelInfo
	NívelDebug
)

// Log representa como deve ser tratado um log na aplicação.
type Log struct {
	outPanic      *log.Logger
	outErro       *log.Logger
	outAviso      *log.Logger
	outInformação *log.Logger
	outDebug      *log.Logger
	Nível         uint
}

// Panic é o método que um log quando acontece um panic na aplicação.
func (log *Log) Panic(imprimir ...any) {
	log.outPanic.Println(imprimir...)
	panic(imprimir)
}

// Erro é o método que faz log de um erro na aplicação.
func (log *Log) Erro(imprimir ...any) {
	if log.Nível < NívelErro {
		return
	}

	log.outErro.Println(imprimir...)
}

// Aviso é o método que faz mensagens de aviso a aplicação.
func (log *Log) Aviso(imprimir ...any) {
	if log.Nível < NívelAviso {
		return
	}

	log.outAviso.Println(imprimir...)
}

// Informação é o método que imprimi mensagens de Informações da aplicação.
func (log *Log) Informação(imprimir ...any) {
	if log.Nível < NívelInfo {
		return
	}

	log.outInformação.Println(imprimir...)
}

// Debug é o método que imprimi mensagens de depuração da aplicação.
func (log *Log) Debug(imprimir ...any) {
	if log.Nível < NívelDebug {
		return
	}

	log.outDebug.Println(imprimir...)
}

// NovoLog cria um log para aplicação.
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

// Arquivos representa os arquivos de saída do log.
type Arquivos struct {
	Pessoa         io.Writer
	Curso          io.Writer
	Aluno          io.Writer
	Professor      io.Writer
	Administrativo io.Writer
	Matéria        io.Writer
	Turma          io.Writer
}

// AbrirArquivos abre os arquivos de log.
func AbrirArquivos(defaultDir string) *Arquivos {
	const flags = os.O_APPEND | os.O_CREATE | os.O_WRONLY

	const mode os.FileMode = 0o666

	const extension = ".log"

	pessoa, err := os.OpenFile(filepath.Clean(defaultDir+"Pessoa"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	curso, err := os.OpenFile(filepath.Clean(defaultDir+"Curso"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	aluno, err := os.OpenFile(filepath.Clean(defaultDir+"Aluno"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	professor, err := os.OpenFile(filepath.Clean(defaultDir+"Professor"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	administrativo, err := os.OpenFile(filepath.Clean(defaultDir+"Administrativo"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	matéria, err := os.OpenFile(filepath.Clean(defaultDir+"Matéria"+extension), flags, mode)
	if err != nil {
		panic(erros.ErroExterno(err))
	}

	turma, err := os.OpenFile(filepath.Clean(defaultDir+"Turma"+extension), flags, mode)
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

// Entidades representa o log de cada entidade da aplicação.
type Entidades struct {
	Pessoa         *Log
	Curso          *Log
	Aluno          *Log
	Professor      *Log
	Administrativo *Log
	Matéria        *Log
	Turma          *Log
}

// NovoLogEntidades cria um log para cada entidade da aplicação.
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
