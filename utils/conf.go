package utils

import (
	"os"

	"github.com/spf13/viper"
)

func GetConf(filename, path string) *viper.Viper {
	dir := os.Getenv("GOPATH") + "\\src\\gos"
	conf := viper.New()
	conf.SetConfigName(filename)
	conf.SetConfigType("yml")
	conf.AddConfigPath(dir + path)
	Panic(conf.ReadInConfig())
	return conf
}
