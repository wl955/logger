package log

import (
	"testing"
)

func init() {
	Init("example", Feishu("qwe"))
}

func TestError(t *testing.T) {
	With("q", 1).With("w", 2).Warnf("%s=%d", "e", 5)
	Warnf("%s=%d", "e", 5)
}
