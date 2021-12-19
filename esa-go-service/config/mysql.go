package config

type Mysql struct {
	Path         string `mapstructure:"path"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Config       string `mapstructure:"config"`
	LogMode      string `mapstructure:"log-mode"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
}
