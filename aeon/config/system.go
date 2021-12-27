package config

type System struct {
	ESAEnv     string `mapstructure:"esa-env"`
	Server     Server `mapstructure:"server"`
	Zap        Zap    `mapstructure:"zap"`
	MysqlESA   Mysql  `mapstructure:"mysql-esa"`
	MysqlBasic Mysql  `mapstructure:"mysql-basic"`
	Redis      Redis  `mapstructure:"redis"`
	KafkaIn    Kafka  `mapstructure:"kafka-in"`
	KafkaOut   Kafka  `mapstructure:"kafka-out"`
	Casbin     Casbin `mapstructure:"casbin"`
	JWT        JWT    `mapstructure:"jwt"`
}
