package logger

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

func init() {
	NewLogrus()
}

// Logrus -
func Logrus() *logrus.Logger {
	return l
}

// Fatal -
func Fatal(args ...interface{}) {
	l.Fatal(args...)
}

// Fatalf - format logger with fatal
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}

// Warn -
func Warn(args ...interface{}) {
	l.Warn(args...)
}

// Warnf -
func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}

// Info -
func Info(args ...interface{}) {
	l.Info(args...)
}

// Infof -
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}

// Error -
func Error(args ...interface{}) {
	l.Error(args...)
}

// Errorf -
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

// Debug -
func Debug(args ...interface{}) {
	l.Debug(args...)
}

// Debugf -
func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}

// Trace -
func Trace(args ...interface{}) {
	l.Trace(args...)
}

// Tracef -
func Tracef(format string, args ...interface{}) {
	l.Tracef(format, args...)
}

// Panic -
func Panic(args ...interface{}) {
	l.Panic(args...)
}

// Panicf -
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}

// WithFields -
func WithFields(f Fields) *logrus.Entry {
	return l.WithFields(f)
}

// WithField -
func WithField(key string, value interface{}) *logrus.Entry {
	return l.WithField(key, value)
}

// WithError -
func WithError(err error) *logrus.Entry {
	return l.WithError(err)
}

// WithContext -
func WithContext(ctx context.Context) *logrus.Entry {
	return l.WithContext(ctx)
}

// Writer -
func Writer() *io.PipeWriter {
	return l.Writer()
}
