package main

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
)

const defaultPort = 80

type Config struct {
	Port    int
	Version int
}

type ConfigBuilder struct {
	port    *int
	version *int
}

func main() {
	cfgBuilder := ConfigBuilder{}
	cfg, err := cfgBuilder.Version(0).Port(10).Build() // Обратная совместимость, можем не указывать необязательные поля
	if err != nil {
		panic(err)
	}
	spew.Dump(cfg)
	NewServer("mandatory", cfg) //Но обязательные прямо в функцию.
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Version(ver int) *ConfigBuilder {
	b.version = &ver
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	var cfg Config
	err := b.validatePort(&cfg)

	return cfg, err
}

func (b *ConfigBuilder) validatePort(cfg *Config) error {
	if b.port == nil {
		cfg.Port = defaultPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return errors.New("negative int")
		} else {
			cfg.Port = *b.port
		}

	}

	return nil
}

func randomPort() int {
	return 1337
}

func NewServer(host string, cfg Config) {

}
