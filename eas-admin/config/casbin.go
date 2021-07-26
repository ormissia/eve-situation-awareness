package config

type Casbin struct {
	ModelPath string `mapstructure:"model-path"` //存放casbin模型的相对路径
}
