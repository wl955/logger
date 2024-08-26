package logfile

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var Writer *rotatelogs.RotateLogs

func init() {
	var e error
	Writer, e = rotatelogs.New("logs/rotatelogs.log.%Y%m%d%H%M",
		rotatelogs.WithLinkName("logs/rotatelogs.log"),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if e != nil {
		panic(e)
	}
}
