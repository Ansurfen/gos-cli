package utils

import (
	"os"

	"github.com/spf13/viper"
)

func GetConf(filename, path string) *viper.Viper {
	return newConf(filename, "yaml", os.Getenv("GOPATH")+"\\src\\gos"+path)
}

func newConf(confName, confType, dir string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigName(confName)
	conf.SetConfigType(confType)
	conf.AddConfigPath(dir)
	Panic(conf.ReadInConfig())
	return conf
}
