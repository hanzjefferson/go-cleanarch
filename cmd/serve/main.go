package main

import (
	"github.com/hanzjefferson/go-cleanarch/internal/bootstrap"
)

func main() {
	viper := bootstrap.NewViper()
	logrus := bootstrap.NewLogrus(viper)
	validator := bootstrap.NewValidator()
	fiber := bootstrap.NewFiber(validator, viper)

	b := bootstrap.Bootstrap{
		Viper:  viper,
		Logrus: logrus,
		Fiber:  fiber,
	}
	b.Boot()
}
