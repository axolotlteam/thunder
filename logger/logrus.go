package logger

import (
	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

// Fatal -
func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

// Warn -
func Warn(args ...interface{}) {
	l.Warn(args...)
}

// Info -
func Info(args ...interface{}) {
	l.Info(args...)
}

// Error -
func Error(args ...interface{}) {
	l.Error(args...)
}

// Debug -
func Debug(args ...interface{}) {
	l.Debug(args...)
}

// Trace -
func Trace(args ...interface{}) {
	l.Trace(args...)
}

// Panic -
func Panic(args ...interface{}) {
	l.Panic(args...)
}

// WithFields -
func WithFields(f Fields) *logrus.Entry {
	return l.WithFields(f)
}
