package database

import (
	"context"
	"testing"
)

func TestNewRedis(t *testing.T) {
	rds := NewRedis("localhost:6379")
	s, err := rds.Ping(context.Background()).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}
