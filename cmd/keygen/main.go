package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/hanzjefferson/go-cleanarch/internal/bootstrap"
)

var generatedSecretKeys = []string{
	"jwt",
}

func main(){
	if len(generatedSecretKeys) == 0 && len(os.Args) < 2 {
		fmt.Printf("there are no secrets that need to be generated.")
	}

	viper := bootstrap.NewViper()
	logrus := bootstrap.NewLogrus(viper)

	keys := generatedSecretKeys
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