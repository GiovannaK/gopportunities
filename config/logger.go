package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		debug:  log.New(writer, "DEBUG: ", logger.Flags()),
		info:   log.New(writer, "INFO: ", logger.Flags()),
		warn:   log.New(writer, "WARN: ", logger.Flags()),
		err:    log.New(writer, "ERROR: ", logger.Flags()),
		writer: writer,
	}
}

// Non formatted log messages
// this is a method of the Logger struct
func (l *Logger) Debug(v ...interface{}) { // empty interface is a type that can hold values of any type
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Formatted log messages

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

