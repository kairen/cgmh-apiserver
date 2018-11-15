package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	expected := "Hello world!!"
	v := GetEnv("GET_ENV1", "Hello world!!")
	assert.Equal(t, expected, v)

	os.Setenv("GET_ENV2", "OS set")
	expected = "OS set"
	v = GetEnv("GET_ENV2", "Test123")
	assert.Equal(t, expected, v)
}
