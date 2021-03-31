package zeimi_raft

import (
	"fmt"
	"io"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warning(v ...interface{})
	Warningf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
}

type raftLogger struct {
	file   *os.File
	writer io.Writer
}

func NewRaftLogger() *raftLogger {
	l := raftLogger{}
	logFile, err := os.OpenFile("./LogFile", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic("can't open Raft Log File:" + string(err.Error()))
	}
	w := io.Writer(logFile)
	l.file = logFile
	l.writer = w
	return &l
}

var GlobalLogger raftLogger

func (r *raftLogger) Debug(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Debugf(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}

func (r *raftLogger) Error(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Errorf(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}
func (r *raftLogger) Info(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Infof(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}

func (r *raftLogger) Warning(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Warningf(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}

func (r *raftLogger) Panic(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Panicf(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}

func (r *raftLogger) Fatal(v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintln(v)))
}

func (r *raftLogger) Fatalf(format string, v ...interface{}) {
	r.writer.Write([]byte("Debug: " + fmt.Sprintf(format, v)))
}
