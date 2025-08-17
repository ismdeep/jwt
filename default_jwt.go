package jwt

import "time"

var defaultJWTClient *JWT

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

// ExpiredAt 获取到期时间
func ExpiredAt(tokens string) (*time.Time, error) {
	return defaultJWTClient.ExpiredAt(tokens)
}
