package log

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	w, e := rotatelogs.New("logs/rotatelogs.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("logs/rotatelogs.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if e != nil {
		panic(e)
	}

	writer = io.MultiWriter(os.Stdout, w)

	wrapper.logger = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(newEncoderConfig()),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(w)),
			zap.NewAtomicLevel(),
		),
		zap.AddStacktrace(zap.NewAtomicLevelAt(zap.ErrorLevel)),
		zap.AddCaller(),
	).Sugar()
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

var wrapper Wrap

type Wrap struct {
	logger *zap.SugaredLogger
}
