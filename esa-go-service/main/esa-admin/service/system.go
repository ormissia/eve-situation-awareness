package service

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"esa-go-service/config"
	"esa-go-service/global"
	model2 "esa-go-service/main/esa-admin/model"
	"esa-go-service/main/esa-admin/model/request"
	source2 "esa-go-service/main/esa-admin/source"
)

func InitDB(conf request.InitDB) (err error) {
	if conf.Path == "" {
		conf.Path = "127.0.0.1:3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", conf.Username, conf.Password, conf.Path)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	myMysqlConfig := config.Mysql{
		Path:     conf.Path,
		DBName:   conf.DBName,
		Username: conf.Username,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	linkDns := myMysqlConfig.Username + ":" + myMysqlConfig.Password + "@tcp(" + myMysqlConfig.Path + ")/" + myMysqlConfig.DBName + "?" + myMysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(myMysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(myMysqlConfig.MaxOpenConns)
		global.ESAMySql = db
	}

	// 初始化表结构
	if err = global.ESAMySql.AutoMigrate(
		model2.Role{},
		model2.User{},
	); err != nil {
		global.ESAMySql = nil
		return err
	}

	// 初始化表数据
	if err = initData(
		source2.Role,
		source2.User,
		source2.UserRole,
		source2.Casbin,
	); err != nil {
		global.ESAMySql = nil
		return err
	}

	return nil
}

// 创建表
func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// 初始化数据
func initData(InitDBFunctions ...request.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
