package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"

	"admin/global"
	"admin/utils"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.EASConfig.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.EASConfig.Zap.Director)
		_ = os.Mkdir(global.EASConfig.Zap.Director, os.ModePerm)
	}

	// 初始化配置文件的Level
	switch global.EASConfig.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	//if level == zap.DebugLevel || level == zap.ErrorLevel {
	//	logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	//} else {
	//	logger = zap.New(getEncoderCore())
	//}
	//if global.GVA_CONFIG.Zap.ShowLine {
	//	logger = logger.WithOptions(zap.AddCaller())
	//}
	return
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.EASConfig.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
