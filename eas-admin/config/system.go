package config

type System struct {
	Server Server `mapstructure:"server"`
	Zap    Zap    `mapstructure:"zap"`
	Mysql  Mysql  `mapstructure:"mysql"`
}
