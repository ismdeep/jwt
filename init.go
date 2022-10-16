package jwt

import (
	"github.com/google/uuid"
	"github.com/ismdeep/config"
)

func init() {
	var data Config
	if err := config.Load("jwt", &data); err == nil {
		Init(&data)
		return
	}

	var ConfigData struct {
		JWT Config
	}
	if err := config.Load("config", &ConfigData); err == nil {
		Init(&data)
		return
	}

	clients = make(map[string]*JWT)
	defaultJWTClient = &JWT{}
	Init(&Config{
		Key:    uuid.NewString(),
		Expire: "72h",
	})
}
