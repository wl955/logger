package log

import (
	"io"

	"go.uber.org/zap"
)

func Writer() io.Writer {
	return writer
}

func Logger() *zap.SugaredLogger {
	return wrapper.logger
}

func Info(args ...interface{}) {
	wrapper.logger.Info(args...)
}

func Warn(args ...interface{}) {
	wrapper.logger.Warn(args...)
}

func Error(args ...interface{}) {
	wrapper.logger.Error(args...)
}

func Panic(args ...interface{}) {
	wrapper.logger.Error(args...)
}

func Infof(template string, args ...interface{}) {
	wrapper.logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	wrapper.logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	wrapper.logger.Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	wrapper.logger.Panicf(template, args...)
}

func With(args ...interface{}) *Wrap {
	return &Wrap{logger: wrapper.logger.With(args...)}
}

func Sync() error {
	return wrapper.logger.Sync()
}
