package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/ismdeep/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {

	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "",
			config: &Config{
				Key:    rand.Password(64, 10, 0),
				Expire: "24h",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.config)
		})
	}
}

func TestCreateToken(t *testing.T) {
	config := &Config{
		Key:    rand.Password(64, 10, 0),
		Expire: "24h",
	}
	Init(config)

	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				content: "user001W",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.content)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestParseToken(t *testing.T) {
	config := &Config{
		Key:    rand.Password(64, 10, 0),
		Expire: "24h",
	}
	Init(config)

	token, err := GenerateToken("user001")
	if err != nil {
		panic(err)
	}

	type args struct {
		tokens string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				tokens: token,
			},
			want:    "user001",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyToken(tt.args.tokens)
			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestCreateTokenWithUserStruct(t *testing.T) {
	user := &struct {
		Username string `json:"username"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	}{
		Username: "ismdeep",
		Nickname: "L. Jiang",
		Avatar:   "https://ismdeep.com/favicon.ico",
	}

	jsonBytes, err := json.Marshal(user)
	assert.NoError(t, err)

	token, err := GenerateToken(string(jsonBytes))
	assert.NoError(t, err)

	fmt.Printf("token = %v\n", token)
}
