package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJwtKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateJwtKey()
			t.Logf("GenerateJwtKey() = %v", got)
		})
	}
}

func TestExpire72Hours(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "72h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire72Hours(), "Expire72Hours()")
		})
	}
}

func TestExpire1Month(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "720h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire1Month(), "Expire1Month()")
		})
	}
}

func TestExpire3Days(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "72h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire3Days(), "Expire3Days()")
		})
	}
}

func TestExpire1Week(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "168h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire1Week(), "Expire1Week()")
		})
	}
}

func TestExpire5Minutes(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "5m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire5Minutes(), "Expire5Minutes()")
		})
	}
}

func TestExpire2Hours(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "2h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire2Hours(), "Expire2Hours()")
		})
	}
}

func TestExpire1Hour(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "1h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire1Hour(), "Expire1Hour()")
		})
	}
}

func TestExpire1Year(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "8760h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire1Year(), "Expire1Year()")
		})
	}
}

func TestExpire100Years(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "876000h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Expire100Years(), "Expire100Years()")
		})
	}
}
