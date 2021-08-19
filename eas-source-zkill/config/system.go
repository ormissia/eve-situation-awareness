package config

type System struct {
	EASEnv string `mapstructure:"eas-env"`
	Server Server `mapstructure:"server"`
	Zap    Zap    `mapstructure:"zap"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
}
