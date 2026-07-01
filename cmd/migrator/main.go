package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/hanzjefferson/go-cleanarch/internal/bootstrap"
	"github.com/pressly/goose/v3"
)

var (
	migrationDir = "./migrations"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  migrator up")
	fmt.Println("  migrator up-by-one")
	fmt.Println("  migrator up-to <version>")
	fmt.Println("  migrator down")
	fmt.Println("  migrator down-to <version>")
	fmt.Println("  migrator redo")
	fmt.Println("  migrator reset")
	fmt.Println("  migrator create <name>")
	fmt.Println("  migrator version")
	fmt.Println("  migrator status")
	os.Exit(1)
}

func parseVersion() (int64, error) {
	if len(os.Args) < 3 {
		return -1, errors.New("missing migration version")
	}

	version, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		return -1, err
	}

	return version, err
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	viper := bootstrap.NewViper()
	logrus := bootstrap.NewLogrus(viper)
	sql := bootstrap.NewSQLDB(viper).DB
	defer sql.Close()

	if err := goose.SetDialect(viper.GetString("database.driver")); err != nil {
		logrus.Fatal(err)
	}

	var err error
	switch os.Args[1] {
	case "up":
		err = goose.Up(sql, migrationDir)

	case "up-by-one":
		err = goose.UpByOne(sql, migrationDir)

	case "up-to":
		ver, err := parseVersion();
		if err != nil {
			logrus.Fatal(err)
		}
		err = goose.UpTo(sql, migrationDir, ver)

	case "down":
		err = goose.Down(sql, migrationDir)

	case "down-to":
		ver, err := parseVersion();
		if err != nil {
			logrus.Fatal(err)
		}
		err = goose.DownTo(sql, migrationDir, ver)

	case "redo":
		err = goose.Redo(sql, migrationDir)

	case "reset":
		err = goose.Reset(sql, migrationDir)

	case "create":
		if len(os.Args) < 3 {
			logrus.Fatal("missing migration name")
		}

		migrationType := "sql"
		if err := goose.Create(nil, migrationDir, os.Args[2], migrationType); err != nil {
			logrus.Fatal(err)
		}

		fmt.Println("Migration created.")
		return

	case "status":
		err = goose.Status(sql, migrationDir)

	case "version":
		version, err := goose.GetDBVersion(sql)
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(version)
		return

	default:
		usage()
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
