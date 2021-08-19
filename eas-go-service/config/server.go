package config

type Server struct {
	Port          int    `mapstructure:"port"`
	TimeFormat    string `mapstructure:"time-format"`
	UseMultipoint bool   `mapstructure:"use-multipoint"` // 多点登录拦截
}
