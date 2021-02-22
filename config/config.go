package config

type Config struct {
	ServerKey          string `yaml:"serverKey"`
	Endpoint           string `yaml:"endpoint"`
	AuthorizedKeysFile string `yaml:"authorizedKeysFile"`
}
