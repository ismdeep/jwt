package jwt

import (
	"errors"
	"testing"

	"bou.ke/monkey"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJWT_VerifyToken(t *testing.T) {
	j := New(&Config{
		Key:    uuid.NewString(),
		Expire: "72h",
	})
	token, err := j.GenerateToken("hello")
	assert.NoError(t, err)
	tokenValid, err := jwt.Parse(token, j.secret())
	assert.NoError(t, err)

	type args struct {
		parseFunc func(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				parseFunc: func(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
					return nil, errors.New("parse failed")
				},
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				parseFunc: func(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
					return &jwt.Token{
						Claims: nil,
					}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "",
			args: args{
				parseFunc: func(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error) {
					return &jwt.Token{
						Claims: tokenValid.Claims,
						Valid:  false,
					}, nil
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monkey.Patch(jwt.Parse, tt.args.parseFunc)

			j := New(&Config{
				Key:    uuid.NewString(),
				Expire: "72h",
			})
			token, err := j.GenerateToken("hello")
			assert.NoError(t, err)
			_, err = j.VerifyToken(token)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
