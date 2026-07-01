package bootstrap

import (
	"time"

	"github.com/hanzjefferson/go-cleanarch/pkg/jwt"
	"github.com/spf13/viper"
)



func NewJWTProvider(config *viper.Viper) *jwt.Provider {
	return &jwt.Provider{
		Issuer: config.GetString("app.name"),
		Secret: []byte(config.GetString("app.secret.jwt")),
		ExpiredDuration: time.Duration(3 * 24 * time.Hour),
	}
}