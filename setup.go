package log

import (
	"io"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

var writer io.Writer

func init() {
	w, e := rotatelogs.New(
		"logs/rotatelogs.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("logs/rotatelogs.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if e != nil {
		log.Fatal(e.Error())
	}

	writer = io.MultiWriter(os.Stdout, w)

	logger = zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(newEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(w)),
			zap.NewAtomicLevel(),
		),
		zap.AddCaller(),
		zap.AddStacktrace(zap.NewAtomicLevelAt(zap.ErrorLevel)),
		zap.AddCallerSkip(1),
	).Sugar()
}

func Writer() io.Writer {
	return writer
}

func Logger() *zap.SugaredLogger {
	return logger
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
