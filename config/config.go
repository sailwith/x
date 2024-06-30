package config

import "github.com/spf13/viper"

type Config struct{}

func New() *Config {
	return &Config{}
}

func (c *Config) Load(file string) error {
	viper.SetConfigFile(file)
	return viper.ReadInConfig()
}

func (c *Config) Unmarshal(data any) error {
	return viper.Unmarshal(data)
}
