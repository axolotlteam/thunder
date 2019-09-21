package logger

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
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

// SetServiceInfo -
func SetServiceInfo(service string) {
	l.AddHook(&fieldHook{
		Service: service,
		Host: func() string {
			name, err := os.Hostname()
			if err != nil {
				return ""
			}
			return name
		}(),
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

// SetSlackHook -
func SetSlackHook(webhook string, level logrus.Level) {

}
