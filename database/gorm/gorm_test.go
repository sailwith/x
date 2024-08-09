package gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := New(Config{
		DSN: dsn,
	})
	assert.NoError(t, err)

	t.Log(db.Name())
}
