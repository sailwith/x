package logger

import "testing"

func TestInfo(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Error(err)
	}

	logger.Info("hello world")
}
