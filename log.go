package log

import (
	"io"
	"os"
	"time"

	"github.com/wlbwlbwlb/log/feishu"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var opt Options

func Init(serviceName string, opts ...OptionFunc) (*zap.SugaredLogger, error) {
	opt = Options{
		service: serviceName,
	}
	for _, fn := range opts {
		fn(&opt)
	}

	w, e := rotatelogs.New("logs/rotatelogs.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("logs/rotatelogs.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if e != nil {
		return logger, e
	}

	encoder := zapcore.NewJSONEncoder(newEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(w)), zapcore.InfoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(feishu.Writer), zapcore.ErrorLevel),
	)

	writer = io.MultiWriter(os.Stdout, w)

	logger = zap.New(
		core,
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	).Sugar()

	return logger, e
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
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

var writer io.Writer

func Writer() io.Writer {
	return writer
}

var logger *zap.SugaredLogger

func Logger() *zap.SugaredLogger {
	return logger
}
