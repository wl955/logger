package log

import (
	"io"

	"go.uber.org/zap"
)

var writer io.Writer

func Writer() io.Writer {
	return writer
}

var logger *zap.SugaredLogger

func Logger() *zap.SugaredLogger {
	return logger
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Panic(args ...interface{}) {
	logger.Error(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func With(args ...interface{}) *Wrap {
	return &Wrap{logger: logger.With(args...)}
}

func Sync() error {
	return logger.Sync()
}
