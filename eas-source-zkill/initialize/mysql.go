package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"eas-source-zkill/global"
)

func Mysql() (db *gorm.DB) {
	m := global.EASConfig.Mysql
	if m.DBName == "" {
		global.EASLog.Warn("Mysql connect failed: need DBName conf")
		return nil
	}

	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.DBName + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.EASLog.Error("Mysql connect failed:", zap.String("err:", err.Error()))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.EASLog.Info("Mysql connect Successful!")
		return db
	}
}

func gormConfig() (config *gorm.Config) {
	config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.EASConfig.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = logger.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = logger.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = logger.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = logger.Default.LogMode(logger.Info)
	default:
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	return
}
