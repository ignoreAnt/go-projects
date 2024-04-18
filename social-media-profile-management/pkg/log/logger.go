package log

import (
	"log"
	"os"
	"sync"
)

// Level LogLevel type to define log levels
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// Logger struct holds the actual logger and the current log level
type Logger struct {
	*log.Logger
	level Level
}

var logger *Logger
var once sync.Once

// GetLogger returns a singleton instance of the Logger
func GetLogger() *Logger {
	once.Do(func() {
		// Get log file name from environment variable
		logFileName := os.Getenv("LOG_FILE")
		if logFileName == "" {
			logFileName = "default.log" // Default file name if none specified
		}

		// Open the log file
		file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("error opening log file: %v", err)
		}

		logger = &Logger{
			Logger: log.New(file, "SMPM <--> ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
			level:  DebugLevel, // Default level; can be set to something else as needed
		}
	})
	return logger
}

// SetLevel sets the logger's level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// Debug logs a debug message if the level is set to Debug
func (l *Logger) Debug(msg string) {
	if l.level <= DebugLevel {
		err := l.Output(2, "DEBUG: "+msg)
		if err != nil {
			return
		} // 2 means count the stack frames to correctly identify the caller
	}
}

// Info logs an informational message if the level is set to Info or lower
func (l *Logger) Info(args ...any) {
	if l.level <= InfoLevel {
		l.SetPrefix("INFO: ")
		l.Println(args...)
	}
}

// Warn logs a warning message if the level is set to Warn or lower
func (l *Logger) Warn(msg string) {
	if l.level <= WarnLevel {
		err := l.Output(2, "WARN: "+msg)
		if err != nil {
			return
		}
	}
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	err := l.Output(2, "ERROR: "+msg)
	if err != nil {
		return
	}
}
