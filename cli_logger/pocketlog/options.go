package pocketlog

import "io"

// Option defines a functional option to our Logger
type Option func(*Logger)

// WithOutput returns the configuration function that sets the output of logs
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}
