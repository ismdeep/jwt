package jwt

import (
	"errors"
	"fmt"
)

var clients map[string]*JWT

func init() {
	clients = make(map[string]*JWT)
}

// InitClient init client
func InitClient(name string, config *Config) error {
	if config == nil || name == "" {
		return errors.New("bad request")
	}
	instance, err := New(config)
	if err != nil {
		return err
	}

	clients[name] = instance
	return nil
}

// InitClients init clients
func InitClients(configs map[string]*Config) error {
	for name, config := range configs {
		if err := InitClient(name, config); err != nil {
			return fmt.Errorf("erorr on [%v], err = [%v]", name, err.Error())
		}
	}
	return nil
}

// GetClient get client
func GetClient(name string) *JWT {
	client, ok := clients[name]
	if !ok {
		return nil
	}

	return client
}
