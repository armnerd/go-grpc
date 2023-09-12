package base

import (
	"encoding/json"
	"os"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var LoggerProvider = wire.NewSet(NewLogger)

type Logger struct {
	logrus *logrus.Logger
}

func NewLogger() Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	return Logger{logger}
}

func (l *Logger) Debug(tag string, info interface{}) {
	l.console(logrus.DebugLevel, tag, info)
}

func (l *Logger) Info(tag string, info interface{}) {
	l.console(logrus.InfoLevel, tag, info)
}

func (l *Logger) Warn(tag string, info interface{}) {
	l.console(logrus.WarnLevel, tag, info)
}

func (l *Logger) Fatal(tag string, info interface{}) {
	l.console(logrus.FatalLevel, tag, info)
}

func (l *Logger) Error(tag string, info interface{}) {
	l.console(logrus.ErrorLevel, tag, info)
}

func (l *Logger) Panic(tag string, info interface{}) {
	l.console(logrus.PanicLevel, tag, info)
}

func (l *Logger) console(level logrus.Level, tag string, something interface{}) {
	// 准备内容
	var info string
	switch value := something.(type) {
	case string:
		info = value
	case map[string]interface{}:
		jsonStr, err := json.Marshal(value)
		if err != nil {
			return
		}
		info = string(jsonStr)
	default:
		return
	}
	// 输出到终端
	field := logrus.Fields{
		"tag": tag,
	}
	switch level {
	case logrus.DebugLevel:
		l.logrus.WithFields(field).Debug(info)
	case logrus.InfoLevel:
		l.logrus.WithFields(field).Info(info)
	case logrus.WarnLevel:
		l.logrus.WithFields(field).Warn(info)
	case logrus.FatalLevel:
		l.logrus.WithFields(field).Fatal(info)
	case logrus.ErrorLevel:
		l.logrus.WithFields(field).Error(info)
	case logrus.PanicLevel:
		l.logrus.WithFields(field).Panic(info)
	}
}
