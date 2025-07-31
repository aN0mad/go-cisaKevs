package cisakev

import "log"

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// defaultLogger implements the Logger interface
type defaultLogger struct{}

func (l *defaultLogger) Debug(msg string, args ...any) { log.Println("DEBUG:", msg) }
func (l *defaultLogger) Info(msg string, args ...any)  { log.Println("INFO:", msg) }
func (l *defaultLogger) Warn(msg string, args ...any)  { log.Println("WARN:", msg) }
func (l *defaultLogger) Error(msg string, args ...any) { log.Println("ERROR:", msg) }
