package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT JWT
type JWT struct {
	signKey    []byte        // 密钥
	clientID   string        // 客户端ID
	ExpireTime time.Duration // 过期时间
}

// AccessClaims jwt claims
type AccessClaims struct {
	Content string `json:"content"`
	jwt.StandardClaims
}

// Init 初始化
func (receiver *JWT) Init(key string, clientID string, expireTimeStr string) {
	var err error

	receiver.signKey = []byte(key)
	receiver.clientID = clientID
	receiver.ExpireTime, err = time.ParseDuration(expireTimeStr)
	if err != nil {
		panic(err)
	}
}

// Valid claims verification
func (a *AccessClaims) Valid() error {
	if time.Unix(a.ExpiresAt, 0).Before(time.Now()) {
		return errors.New("invalid access token")
	}
	return nil
}

// CreateToken 生成token
func (receiver *JWT) CreateToken(content string) (token string, err error) {
	claim := &AccessClaims{
		Content: content,
		StandardClaims: jwt.StandardClaims{
			Audience:  receiver.clientID,
			ExpiresAt: time.Now().Add(receiver.ExpireTime).Unix(),
			Subject:   receiver.clientID,
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err = tokens.SignedString(receiver.signKey)
	return token, err
}

// ParseToken 格式化token
func (receiver *JWT) ParseToken(tokens string) (string, error) {
	token, err := jwt.Parse(tokens, receiver.secret())
	if err != nil {
		return "", err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to map claim")
		return "", err
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return "", err
	}
	content := claim["content"].(string)
	return content, nil
}

func (receiver *JWT) secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return receiver.signKey, nil
	}
}
