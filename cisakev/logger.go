package cisakev

import "log"

type Logger interface {
	Trace(msg string, args ...any)
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
}

// defaultLogger implements the Logger interface
type defaultLogger struct{}

func (l *defaultLogger) Trace(msg string, args ...any) { log.Println("TRACE:", msg) }
func (l *defaultLogger) Debug(msg string, args ...any) { log.Println("DEBUG:", msg) }
func (l *defaultLogger) Info(msg string, args ...any)  { log.Println("INFO:", msg) }
func (l *defaultLogger) Warn(msg string, args ...any)  { log.Println("WARN:", msg) }
func (l *defaultLogger) Error(msg string, args ...any) { log.Println("ERROR:", msg) }
func (l *defaultLogger) Fatal(msg string, args ...any) { log.Fatalln("FATAL:", msg) }
