package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedis(t *testing.T) {
	rds := NewRedis("localhost:6379")
	s, err := rds.Ping(context.Background()).Result()
	assert.NoError(t, err)
	t.Log(s)
}
