package log

import (
	"fmt"
	"io"
	"log"
)

var l Logger

type Logger struct {
	*log.Logger
}

func InitLogger(writer io.Writer) {
	l = Logger{log.New(writer, "", log.LstdFlags)}
}

func Debug(v ...interface{}) {
	l.Print("- debug - ", fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	l.Logger.Printf(fmt.Sprintf("- debug - %s", format), v...)
}

func Info(v ...interface{}) {
	l.Print("- info - ", fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	l.Logger.Printf(fmt.Sprintf("- info - %s", format), v...)
}

func Error(v ...interface{}) {
	l.Print("- error - ", fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	l.Logger.Printf(fmt.Sprintf("- error - %s", format), v...)
}
