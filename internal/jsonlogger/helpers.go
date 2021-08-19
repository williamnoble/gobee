package jsonlogger

import "os"

func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.printTemplate(LevelInfo, message, properties)
}

// PrintError writes a standard error to given writer
func (l *Logger) PrintError(err error, properties map[string]string) {
	l.printTemplate(LevelError, err.Error(), properties)
}

// PrintFatal will print Fatal errors and then suspend exucution via calling os.Exit(1)
func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.printTemplate(LevelFatal, err.Error(), properties)
	os.Exit(1)
}