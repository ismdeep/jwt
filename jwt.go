package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/ismdeep/rand"
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
func New(config *Config) *JWT {
	c := &Config{
		Key:    rand.HexStr(32),
		Expire: "72h",
	}
	if config != nil {
		c = config
	}

	if _, err := time.ParseDuration(c.Expire); err != nil {
		c.Expire = "72h"
	}

	instance := &JWT{}
	instance.signKey = []byte(c.Key)
	instance.expireTime, _ = time.ParseDuration(c.Expire)

	return instance
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
