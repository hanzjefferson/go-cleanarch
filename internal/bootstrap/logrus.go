package bootstrap

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogrus(config *viper.Viper) *logrus.Entry {
	l := logrus.New()

	if (config.GetBool("app.debug")) {
		l.SetLevel(logrus.DebugLevel)
	}

	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logrus.NewEntry(l)
}