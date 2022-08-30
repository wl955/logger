package logger

import (
	"io"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var logger *zap.SugaredLogger

func Setup() (w io.Writer, f func()) {
	w, e := rotatelogs.New(
		"logs/rotatelogs.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("logs/rotatelogs.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if e != nil {
		log.Fatal(e.Error())
	}

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

	return w, func() {
		logger.Sync()
	}
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
