package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitClient(t *testing.T) {
	type args struct {
		name   string
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// OK: normal arguments
			name: "",
			args: args{
				name: "config1",
				config: &Config{
					Key:    "12345678",
					Expire: "24h",
				},
			},
			wantErr: false,
		},
		{
			// ERROR: config is nil
			name: "",
			args: args{
				name:   "config2",
				config: nil,
			},
			wantErr: true,
		},
		{
			// ERROR: name is empty
			name: "",
			args: args{
				name: "",
				config: &Config{
					Key:    "12345678",
					Expire: "24h",
				},
			},
			wantErr: true,
		},
		{
			// ERROR: expire string is invalid
			name: "",
			args: args{
				name: "config3",
				config: &Config{
					Key:    "12345678",
					Expire: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitClient(tt.args.name, tt.args.config)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestInitClients(t *testing.T) {
	type args struct {
		configs map[string]*Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				configs: map[string]*Config{
					"config1": {
						Key:    "12345678",
						Expire: "24h",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				configs: map[string]*Config{
					"config1": {
						Key:    "12345678",
						Expire: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitClients(tt.args.configs)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestGetClient(t *testing.T) {
	err := InitClients(map[string]*Config{
		"config1": {
			Key:    "12345678",
			Expire: "24h",
		},
	})
	assert.NoError(t, err)

	client1 := GetClient("config1")
	assert.NotNil(t, client1)

	client2 := GetClient("config-invalid")
	assert.Nil(t, client2)

}
