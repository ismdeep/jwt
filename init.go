package jwt

import "github.com/ismdeep/rand"

func init() {
	clients = make(map[string]*JWT)
	defaultJWTClient = &JWT{}
	Init(&Config{
		Key:    rand.Password(64, 10, 0),
		Expire: "72h",
	})
}
