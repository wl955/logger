package log

import (
	"go.uber.org/zap"
	"io"
)

func Writer() io.Writer {
	return writer
}

func Logger() *zap.SugaredLogger {
	return L
}
