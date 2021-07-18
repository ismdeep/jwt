package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		key           string
		expireTimeStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				key:           "12345678",
				expireTimeStr: "24h",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.key, tt.args.expireTimeStr); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateToken(t *testing.T) {

	err := Init("12345678", "24h")
	if err != nil {
		panic(err)
	}

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
			got, err := GenerateToken(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("token = %v\n", got)
		})
	}
}

func TestParseToken(t *testing.T) {

	err := Init("12345678", "24h")
	if err != nil {
		panic(err)
	}

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
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyToken() got = %v, want %v", got, tt.want)
			}
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
	if err != nil {
		panic(err)
	}

	token, err := GenerateToken(string(jsonBytes))
	if err != nil {
		panic(err)
	}

	fmt.Printf("token = %v\n", token)
}
