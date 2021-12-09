package config

type Kafka struct {
	Path     string `mapstructure:"path"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Topic    string `mapstructure:"topic"`
	Group    string `mapstructure:"group"`
}
