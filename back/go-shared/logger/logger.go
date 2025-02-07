package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelColors = map[LogLevel]string{
	DEBUG: "\033[36m", // Cyan
	INFO:  "\033[32m", // Green
	WARN:  "\033[33m", // Yellow
	ERROR: "\033[31m", // Red
	FATAL: "\033[35m", // Magenta
}

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) log(level LogLevel, msg string, a ...any) {
	color := levelColors[level]
	reset := "\033[0m"

	_, file, line, _ := runtime.Caller(2)
	fileName := filepath.Base(file)
	levelStr := strings.ToUpper(level.String())

	caller := fmt.Sprintf("%s:%d", fileName, line)
	fmsg := fmt.Sprintf(msg, a...)

	logEntry := fmt.Sprintf("%s[%s][%s] %s%s", color, levelStr, caller, fmsg, reset)

	if level == FATAL {
		log.Fatal(logEntry)
	} else {
		log.Println(logEntry)
	}
}

func (l *Logger) Debug(msg string, a ...any) { l.log(DEBUG, msg, a...) }
func (l *Logger) Info(msg string, a ...any)  { l.log(INFO, msg, a...) }
func (l *Logger) Warn(msg string, a ...any)  { l.log(WARN, msg, a...) }
func (l *Logger) Error(msg string, a ...any) { l.log(ERROR, msg, a...) }
func (l *Logger) Fatal(msg string, a ...any) { l.log(FATAL, msg, a...) }

func (level LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]
}

// Global

var GlobalLogger *Logger

func init() {
	GlobalLogger = NewLogger()
}

func Debug(msg string, a ...any) { GlobalLogger.log(DEBUG, msg, a...) }
func Info(msg string, a ...any)  { GlobalLogger.log(INFO, msg, a...) }
func Warn(msg string, a ...any)  { GlobalLogger.log(WARN, msg, a...) }
func Error(msg string, a ...any) { GlobalLogger.log(ERROR, msg, a...) }
func Fatal(msg string, a ...any) { GlobalLogger.log(FATAL, msg, a...) }
