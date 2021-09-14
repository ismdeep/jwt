package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT Struct Info
type JWT struct {
	signKey    []byte        // 密钥
	expireTime time.Duration // 过期时间
}

// claims jwt claims
type claimsStruct struct {
	Content string `json:"content"`
	jwt.StandardClaims
}

// New create instance
func New(config *Config) (*JWT, error) {
	if config == nil {
		return nil, errors.New("bad request")
	}
	var err error
	instance := &JWT{}
	instance.signKey = []byte(config.Key)
	instance.expireTime, err = time.ParseDuration(config.Expire)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// GenerateToken generate token
func (receiver *JWT) GenerateToken(content string) (token string, err error) {
	claim := &claimsStruct{
		Content: content,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(receiver.expireTime).Unix(),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err = tokens.SignedString(receiver.signKey)
	return token, err
}

// VerifyToken verify token
func (receiver *JWT) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, receiver.secret())
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

	_, ok = claim["content"]
	if !ok {
		return "", errors.New("invalid token")
	}

	content := claim["content"].(string)
	return content, nil
}

func (receiver *JWT) secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return receiver.signKey, nil
	}
}
