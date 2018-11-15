package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	email := "test@gmail.com"
	uuid := "u00001"
	token, err := GenerateToken(email, uuid, time.Hour*1)
	if err != nil {
		t.Fatalf("Could not generate token, %s.", err.Error())
	}

	claims, err := ParseToken(token)
	if err != nil {
		t.Fatalf("Could not parse token, %s.", err.Error())
	}

	expected := &Claims{
		Email: Base64Encode(email),
		UUID:  Base64Encode(uuid),
	}

	assert.Equal(t, expected.Email, claims.Email)
	assert.Equal(t, expected.UUID, claims.UUID)
}
