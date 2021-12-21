package config

type System struct {
	ESAEnv string `mapstructure:"esa-env"`
	Server Server `mapstructure:"server"`
	Zap    Zap    `mapstructure:"zap"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Casbin Casbin `mapstructure:"casbin"`
	JWT    JWT    `mapstructure:"jwt"`
}
