package logs

import (
	"io"
	"log"
)

type Log struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

func NewLog(out io.Writer) *Log {
	return &Log{
		Info:    log.New(out, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(out, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(out, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
