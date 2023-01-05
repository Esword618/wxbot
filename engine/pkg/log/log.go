package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type logger struct {
	l          *logrus.Logger
	callerFile string
	callerLine int
}

var log = &logger{
	l: logrus.New(),
}

type Formatter struct{}

func (s *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG_LOG") == "true" {
		return []byte(fmt.Sprintf("[%s] [%s] [%s:%d] %s\n", timestamp, level, log.callerFile, log.callerLine, entry.Message)), nil
	} else {
		return []byte(fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, entry.Message)), nil
	}
}

func init() {
	log.l.SetLevel(logrus.TraceLevel)
	log.l.SetOutput(os.Stdout)
	//log.l.SetFormatter(&Formatter{})
	log.l.SetReportCaller(true)
	log.l.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			if os.Getenv("DEBUG") == "true" || os.Getenv("DEBUG_LOG") == "true" {
				return "", fmt.Sprintf("[%s:%d]", log.callerFile, log.callerLine)
			}
			return "", ""
		},
	})
}

func GetLogger() *logrus.Logger {
	return log.l
}

func getCaller() {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	log.callerFile = path.Join(path.Base(path.Dir(file)), path.Base(file))
	log.callerLine = line
}

func Println(args ...interface{}) {
	getCaller()
	log.l.Println(args...)
}

func Printf(format string, args ...interface{}) {
	getCaller()
	log.l.Printf(format, args...)
}

func Debug(args ...interface{}) {
	getCaller()
	log.l.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	getCaller()
	log.l.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	getCaller()
	log.l.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	getCaller()
	log.l.Warnf(format, args...)
}

func Error(args ...interface{}) {
	getCaller()
	log.l.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	getCaller()
	log.l.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	getCaller()
	log.l.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	getCaller()
	log.l.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	getCaller()
	log.l.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	getCaller()
	log.l.Panicf(format, args...)
}

func Trace(args ...interface{}) {
	getCaller()
	log.l.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	getCaller()
	log.l.Tracef(format, args...)
}
