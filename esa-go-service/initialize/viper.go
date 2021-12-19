package initialize

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/spf13/viper"

	"esa-go-service/global"
)

func Viper(fileName string) (v *viper.Viper) {
	file, err := global.ESAStaticFile.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v = viper.New()
	v.SetConfigType("yaml")
	err = v.ReadConfig(bytes.NewBuffer(file))
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&global.ESAConfig); err != nil {
		fmt.Println(err)
	}
	return
}
