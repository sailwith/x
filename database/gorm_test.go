package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQL(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := NewMySQL(dsn)
	assert.NoError(t, err)

	t.Log(db.Name())
}
