package config

import (
	"time"
)

type Config struct {
	ServerKey          string        `yaml:"serverKey"`
	Endpoints          []string      `yaml:"endpoints"`
	AuthorizedKeysFile string        `yaml:"authorizedKeysFile"`
	LoopTime           time.Duration `yaml:"loopTime"`
	Tries              int           `yaml:"tries"`
	RetryTimeout       time.Duration `yaml:"retryTimeout"`
}

func DefaultConfig() *Config {
	return &Config{
		LoopTime:     1 * time.Minute,
		RetryTimeout: 1 * time.Second,
		Tries:        1,
	}
}
