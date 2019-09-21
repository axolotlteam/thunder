package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var l *logrus.Logger

// Fields -
type Fields = logrus.Fields

// log level
const (
	TraceLevel = logrus.TraceLevel
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
)

//NewLogrus -
func NewLogrus() {
	l = logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			//return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			return f.Function + "()", filename + ":" + strconv.Itoa(f.Line)
		},
	})
}

// SetLevel -
func SetLevel(level logrus.Level) {
	switch level {
	case TraceLevel:
		l.SetLevel(logrus.TraceLevel)
	case DebugLevel:
		l.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		l.SetLevel(logrus.InfoLevel)
	case ErrorLevel:
		l.SetLevel(logrus.ErrorLevel)
	case WarnLevel:
		l.SetLevel(logrus.WarnLevel)
	}
}

// AddSlackHook -
func AddSlackHook(webhook string) {
}
