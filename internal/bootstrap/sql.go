package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

func NewSQLDB(config *viper.Viper) *sql.DB {
	driver := config.GetString("database.driver")

	var dsn string
	if config.IsSet("database.dsn") {
		dsn = config.GetString("database.dsn")
	} else {
		switch config.GetString("database.driver") {
		case "mysql":
			username := config.GetString("database.username")
			password := config.GetString("database.password")
			host := config.GetString("database.host")
			port := config.GetInt("database.port")
			name := config.GetString("database.name")
			dsn = genMySQLDSN(username, password, host, port, name)
		case "sqlite3":
			name := config.GetString("database.name")
			dsn = genSQLite3DSN(name)
		}
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(fmt.Errorf("database connect error:\n%v", err))
	}
	defer db.Close()

	return db
}

func genMySQLDSN(username string, password string, host string, port int, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, name)
}

func genSQLite3DSN(filename string) string {
	return fmt.Sprintf("file:%s?cache=shared", filename)
}
