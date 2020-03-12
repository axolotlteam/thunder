package logger

// Logger -
type Logger interface {
	Fatal(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
