package main

import (
	"github.com/hanzjefferson/go-cleanarch/internal/bootstrap"
)

func main() {
	viper := bootstrap.NewViper()
	logrus := bootstrap.NewLogrus(viper)
	validator := bootstrap.NewValidator()
	fiber := bootstrap.NewFiber(validator, viper)
	sql := bootstrap.NewSQLDB(viper)
	jwt := bootstrap.NewJWTProvider(viper)

	b := bootstrap.Bootstrap{
		Config:  viper,
		Log: logrus,
		Fiber:  fiber,
		SQL: sql,
		JWT: jwt,
	}
	b.Boot()
}
