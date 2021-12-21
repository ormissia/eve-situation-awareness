package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key"`  //jwt签名
	ExpiresTime int64  `mapstructure:"expires-time"` //过期时间
	BufferTime  int64  `mapstructure:"buffer-time"`  //缓冲时间
}
