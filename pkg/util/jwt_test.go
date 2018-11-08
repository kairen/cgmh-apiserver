package util

import (
	"strings"
	"testing"
	"time"
)

func TestBase64Encode(t *testing.T) {
	expected := "SGVsbG8gV29ybGQhIQ=="
	encoded := Base64Encode("Hello World!!")
	if strings.Compare(expected, encoded) != 0 {
		t.Errorf("Wrong encoded string. Expected %s, got %s", expected, encoded)
	}
}

func TestBase64Decode(t *testing.T) {
	expected := "Hello World!!"
	decoded, err := Base64Decode("SGVsbG8gV29ybGQhIQ==")
	if err != nil {
		t.Fatalf("Could not decode string, %s.", err.Error())
	}

	if strings.Compare(expected, decoded) != 0 {
		t.Errorf("Wrong decoded string. Expected %s, got %s", expected, decoded)
	}
}

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
	if strings.Compare(expected.Email, claims.Email) != 0 {
		t.Errorf("Wrong parse jwt email. Expected %s, got %s", expected.Email, claims.Email)
	}

	if strings.Compare(expected.UUID, claims.UUID) != 0 {
		t.Errorf("Wrong  parse jwt uuid. Expected %s, got %s", expected.UUID, claims.UUID)
	}
}
