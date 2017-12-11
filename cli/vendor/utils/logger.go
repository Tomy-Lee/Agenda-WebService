package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	loggers []*log.Logger
	output  io.WriteCloser
)

// NewLogger returns a new logger with some default settings
func NewLogger() *log.Logger {
	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	loggers = append(loggers, logger)
	return logger
}

// InitLogFile prepares log file
func InitLogFile(logPath string) {
	logFile := openLogFile(logPath)
	setLoggersOutput(logFile)
}

// CloseLogFile closes opened log file
func CloseLogFile() error {
	if output != nil {
		return output.Close()
	}
	return nil
}

func setLoggersOutput(w io.Writer) {
	for _, logger := range loggers {
		logger.SetOutput(w)
	}
}

func openLogFile(path string) io.Writer {
	if len(path) > 0 {
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			output = file
			return file
		}
		fmt.Fprintf(os.Stderr, `Failed to open file "%s" for logging: %v\n`, path, err)
	}
	if _, present := os.LookupEnv("DEBUG"); present {
		return os.Stdout
	}
	return ioutil.Discard
}
