package jwt

var defaultJWTClient *JWT

func init() {
	defaultJWTClient = &JWT{}
}

// Init 初始化
func Init(key string, expireTimeStr string) error {
	return defaultJWTClient.Init(key, expireTimeStr)
}

// GenerateToken 生成token
func GenerateToken(content string) (string, error) {
	return defaultJWTClient.GenerateToken(content)
}

// VerifyToken 格式化token
func VerifyToken(tokens string) (string, error) {
	return defaultJWTClient.VerifyToken(tokens)
}
