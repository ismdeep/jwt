package jwt

var defaultJWTClient *JWT

func init() {
	defaultJWTClient = &JWT{}
}

// Init 初始化
func Init(key string, clientID string, expireTimeStr string) {
	defaultJWTClient.Init(key, clientID, expireTimeStr)
}

// CreateToken 生成token
func CreateToken(content string) (token string, err error) {
	return defaultJWTClient.CreateToken(content)
}

// ParseToken 格式化token
func ParseToken(tokens string) (string, error) {
	return defaultJWTClient.ParseToken(tokens)
}
