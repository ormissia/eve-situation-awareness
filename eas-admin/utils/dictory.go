package utils

import (
	"go.uber.org/zap"
	"os"

	"admin/global"
)

// PathExists 文件目录是否存在
func PathExists(path string) (exist bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.EASLog.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.EASLog.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}
