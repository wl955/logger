package log

import "testing"

func TestError(t *testing.T) {
	With("a", 1).With("b", 2).Warnf("%s,%d", "c", 3)
	Warnf("%s,%d", "c", 3)
}
