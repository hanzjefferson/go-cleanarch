package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/hanzjefferson/go-cleanarch/internal/bootstrap"
)

func main(){
	viper := bootstrap.NewViper()
	logrus := bootstrap.NewLogrus(viper)

	keys := slices.Collect(maps.Keys(viper.GetStringMap("app.secret")))
	if len(os.Args) > 1 {
		keys = os.Args[1:]
	}
	for _, k := range keys {
		logrus.Infof("generating secret for '%s'...", k)
		secret, err := gen()
		if err != nil {
			logrus.Warnf("failed to generate secret for '%s':%+v", k, err)
		}
		logrus.Debugf("secret generated: %s", secret)

		viper.Set(fmt.Sprintf("app.secret.%s", k), secret)
	}

	logrus.Infof("saving config....")
	if err := viper.WriteConfig(); err != nil {
		logrus.Fatalf("failed to save config:%+v", err)
	}
}

func gen() (string, error) {
	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(secret), nil
}