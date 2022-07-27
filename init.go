package jwt

import (
	"github.com/ismdeep/config"
	"github.com/ismdeep/rand"
)

func init() {
	var data Config
	if err := config.Load("jwt", &data); err == nil {
		Init(&data)
		return
	}

	clients = make(map[string]*JWT)
	defaultJWTClient = &JWT{}
	Init(&Config{
		Key:    rand.Password(64, 10, 0),
		Expire: "72h",
	})
}
