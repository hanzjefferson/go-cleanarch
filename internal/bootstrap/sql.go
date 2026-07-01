package bootstrap

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func NewSQLDB(config *viper.Viper) *sqlx.DB {
	driver := "mysql"

	var dsn string
	if config.IsSet("database.dsn") {
		dsn = config.GetString("database.dsn")
	} else {
		//MySQl
		username := config.GetString("database.username")
		password := config.GetString("database.password")
		host := config.GetString("database.host")
		port := config.GetInt("database.port")
		name := config.GetString("database.name")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
		
		// SQLite3
		// name := config.GetString("database.name")
		// dsn = fmt.Sprintf("file:%s?cache=shared", name)
	}

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(fmt.Errorf("database connect error:\n%v", err))
	}

	return db
}
