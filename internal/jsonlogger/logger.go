package jsonlogger

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
	LevelOff
)

type logMetadata struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Props   map[string]string `json:"props"`
	Trace string `json:"debug,omitempty"`
}

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERRROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

type Logger struct {
	logWriter io.Writer
	minLevel  Level
	mu        sync.Mutex // allow concurrent write access
}

func New(writer io.Writer, minimumLogLevel Level) *Logger {
	return &Logger{
		logWriter: os.Stdout,
		minLevel:  minimumLogLevel,
		mu:        sync.Mutex{},
	}
}

func (l *Logger) printTemplate(level Level, msg string, props map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	var debugTrace []byte
	if level > LevelError {  // LevelError: Level 1
		debugTrace = debug.Stack() // []byte
	}


	data := logMetadata{
		Level:   level.String(),
		Time:    time.Now().Format(time.RFC3339), // fmt as String
		Message: msg,
		Props:   props,
		Trace:  string(debugTrace) ,
	}

	var jsonifiedLogRow []byte
	jsonifiedLogRow, err := json.Marshal(data)
	if err != nil {
		jsonifiedLogRow = []byte(LevelError.String() + ": error in marshalling error log" + err.Error())
	}

	jsonifiedLogRow = append(jsonifiedLogRow, '\n')

	l.mu.Lock()
	defer l.mu.Unlock()
	return l.logWriter.Write(jsonifiedLogRow)
}

// A Neccessity for writing logs for http std server.
func (l *Logger) Write(message []byte) (n int, err error) {
	return l.printTemplate(LevelError, string(message), nil)
}

