package config

type Redis struct {
	Path     string `mapstructure:"path"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
