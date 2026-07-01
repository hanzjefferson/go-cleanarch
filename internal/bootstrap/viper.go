package bootstrap

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")

	v.SetDefault("app.debug", true)
	v.SetDefault("app.name", "go-application")
	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", 3000)
	v.SetDefault("database.driver", "mysql")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper couldn't read config:\n%v", err))
	}

	return v
}