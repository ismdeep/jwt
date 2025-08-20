package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// GenerateJwtKey generate jwt key
func GenerateJwtKey() string {
	key := make([]byte, 32) // 256 位密钥
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("failed to generate random key: %v", err)
	}
	return base64.RawURLEncoding.EncodeToString(key)
}

func Expire72Hours() string {
	return "72h"
}

func Expire1Month() string {
	return "720h"
}

func Expire3Days() string {
	return Expire72Hours()
}

func Expire1Week() string {
	return "168h"
}

func Expire5Minutes() string {
	return "5m"
}

func Expire2Hours() string {
	return "2h"
}

func Expire1Hour() string {
	return "1h"
}

func Expire1Year() string {
	return "8760h"
}

func Expire100Years() string {
	return "876000h"
}
