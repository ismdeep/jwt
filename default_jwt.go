package jwt

import (
	"github.com/ismdeep/rand"
)

var defaultJWTClient *JWT

func init() {
	defaultJWTClient = &JWT{}
	Init(&Config{
		Key:    rand.HexStr(128),
		Expire: "72h",
	})
}

// Init 初始化
func Init(config *Config) {
	defaultJWTClient = New(config)
}

// GenerateToken 生成token
func GenerateToken(content string) (string, error) {
	return defaultJWTClient.GenerateToken(content)
}

// VerifyToken 格式化token
func VerifyToken(tokens string) (string, error) {
	return defaultJWTClient.VerifyToken(tokens)
}
