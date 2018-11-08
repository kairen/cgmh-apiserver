package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	UUID  string `json:"uuid"`
	jwt.StandardClaims
}

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
		return string(decode), err
	}
	return string(decode), nil
}

func GenerateToken(email, uuid string, expTime time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expTime)

	claims := Claims{
		Base64Encode(email),
		Base64Encode(uuid),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte("secret"))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
