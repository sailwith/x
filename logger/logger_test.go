package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	logger, err := New()
	assert.NoError(t, err)

	logger.Info("hello world")
}
