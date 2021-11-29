package logs

import (
	"io"
	"log"
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
