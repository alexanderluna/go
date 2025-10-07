package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a logger, ready to log at the required threshold.
// Give it a list of configuration functions to adapt it to your own needs.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}
	// loop over configuration functions and pass it the logger to apply them
	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	l.logf(format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}

	l.logf(format, args...)
}

// Warnf formats and prints a message if the log level is warn or higher.
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}

	l.logf(format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}

	l.logf(format, args...)
}

// Fatalf formats and prints a message if the log level is fatal or higher.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}

	l.logf(format, args...)
}

// logf prints the message to the output
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
