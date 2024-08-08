package log

import "testing"

func TestError(t *testing.T) {
	With("a", 1).With("s", 2).With("d", 3).With("f", 4).Warnf("%s=%d", "g", 5)
	Warnf("%s=%d", "g", 5)
}
