package config

type Server struct {
	Port       int    `mapstructure:"port"`
	TimeFormat string `mapstructure:"time-format"`
}
