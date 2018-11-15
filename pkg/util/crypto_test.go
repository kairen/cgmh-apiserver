package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	expected := "SGVsbG8gV29ybGQhIQ=="
	encoded := Base64Encode("Hello World!!")
	assert.Equal(t, expected, encoded)
}

func TestBase64Decode(t *testing.T) {
	expected := "Hello World!!"
	decoded, err := Base64Decode("SGVsbG8gV29ybGQhIQ==")
	if err != nil {
		t.Fatalf("Could not decode string, %s.", err.Error())
	}
	assert.Equal(t, expected, decoded)
}

func TestMD5Encode(t *testing.T) {
	expected := "cbf41347bb1978f6f32087b2cf01e351"
	encoded := MD5Encode("Hello World!!")
	assert.Equal(t, expected, encoded)
}
