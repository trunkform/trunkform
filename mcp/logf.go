package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	FATAL LogLevel = 1
	ERROR LogLevel = 2
	WARN  LogLevel = 3
	INFO  LogLevel = 4
	DEBUG LogLevel = 5
	TRACE LogLevel = 6
)

type Logger struct {
	level  LogLevel
	writer io.Writer
}

var logf = NewLogger(os.Getenv("TRUNKFORM_LOGLEVEL"), os.Stderr)

func NewLogger(levelStr string, w io.Writer) *Logger {
	levelMap := map[string]LogLevel{
		"1": FATAL, "2": ERROR, "3": WARN,
		"4": INFO, "5": DEBUG, "6": TRACE,
		"FATAL": FATAL, "ERROR": ERROR, "WARN": WARN,
		"INFO": INFO, "DEBUG": DEBUG, "TRACE": TRACE,
	}
	level, ok := levelMap[levelStr]
	if !ok {
		level = INFO
	}
	return &Logger{level: level, writer: w}
}

func (l *Logger) Log(level LogLevel, prefix, message string, colorCode ...string) {
	if level > l.level {
		return
	}
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	color := ""
	if len(colorCode) > 0 {
		color = colorCode[0]
	}
	fmt.Fprintf(l.writer, "%s [%s] %s%s\x1b[0m\n", timestamp, prefix, color, message)
}

func (l *Logger) Fatalf(message string, colorCode ...string)  { l.Log(FATAL, "FATAL", message, colorCode...) }
func (l *Logger) Errorf(message string, colorCode ...string)  { l.Log(ERROR, "ERROR", message, colorCode...) }
func (l *Logger) Warnf(message string, colorCode ...string)   { l.Log(WARN, "WARN", message, colorCode...) }
func (l *Logger) Infof(message string, colorCode ...string)   { l.Log(INFO, "INFO", message, colorCode...) }
func (l *Logger) Debugf(message string, colorCode ...string)  { l.Log(DEBUG, "DEBUG", message, colorCode...) }
func (l *Logger) Tracef(message string, colorCode ...string)  { l.Log(TRACE, "TRACE", message, colorCode...) }
