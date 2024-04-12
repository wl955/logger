package log

import "testing"

func TestError(t *testing.T) {
	L.With("1", 2).With("3", 4).Info("5", "6", 7, 8)
	//{"level":"info","ts":"2024-04-12 10:24:24","caller":"log/logger_test.go:6","msg":"567 8","1":2,"3":4}
	L.With("1", 2).With("3", 4).Infof("%s,%s,%d,%d", "5", "6", 7, 8)
	//{"level":"info","ts":"2024-04-12 10:26:01","caller":"log/logger_test.go:8","msg":"5,6,7,8","1":2,"3":4}
}
