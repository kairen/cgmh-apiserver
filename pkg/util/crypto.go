package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Base64Encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func Base64Decode(v string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}
	return string(decode), nil
}

func MD5Encode(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
