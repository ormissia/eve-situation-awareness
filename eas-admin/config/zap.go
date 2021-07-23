package config

type Zap struct {
	Level         string `mapstructure:"level"`          // 级别
	Format        string `mapstructure:"format"`         // 输出
	Prefix        string `mapstructure:"prefix"`         // 前缀
	Director      string `mapstructure:"director"`       // 日志文件夹
	LinkName      string `mapstructure:"link-name"`      // 软链接名称
	ShowLine      bool   `mapstructure:"show-line"`      // 显示行
	EncodeLevel   string `mapstructure:"encode-level"`   // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console"` // 输出控制台
}
