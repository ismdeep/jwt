package jwt

import (
	"github.com/ismdeep/rand"
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
					Key:    rand.Password(64, 10, 0),
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
					Key:    rand.Password(64, 10, 0),
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
					Key:    rand.Password(64, 10, 0),
					Expire: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitClient(tt.args.name, tt.args.config)
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
						Key:    rand.Password(64, 10, 0),
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
						Key:    rand.Password(64, 10, 0),
						Expire: "",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitClients(tt.args.configs)
		})
	}
}

func TestGetClient(t *testing.T) {
	InitClients(map[string]*Config{
		"config1": {
			Key:    rand.Password(64, 10, 0),
			Expire: "24h",
		},
	})

	client1 := GetClient("config1")
	assert.NotNil(t, client1)

	client2 := GetClient("config-invalid")
	assert.Nil(t, client2)

}
