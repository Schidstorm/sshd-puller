package config

import "time"

type Config struct {
	ServerKey          string        `yaml:"serverKey"`
	Endpoint           string        `yaml:"endpoint"`
	AuthorizedKeysFile string        `yaml:"authorizedKeysFile"`
	LoopTime           time.Duration `yaml:"loopTime"`
}
