package main

import (
	"fmt"
	"log"
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

func parseVersion() int64 {
	if len(os.Args) < 3 {
		log.Fatal("missing migration version")
	}

	version, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("invalid version: %v", err)
	}

	return version
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	viper := bootstrap.NewViper()
	sql := bootstrap.NewSQLDB(viper).DB
	defer sql.Close()

	if err := goose.SetDialect(viper.GetString("database.driver")); err != nil {
		log.Fatal(err)
	}

	var err error
	switch os.Args[1] {
	case "up":
		err = goose.Up(sql, migrationDir)

	case "up-by-one":
		err = goose.UpByOne(sql, migrationDir)

	case "up-to":
		err = goose.UpTo(sql, migrationDir, parseVersion())

	case "down":
		err = goose.Down(sql, migrationDir)

	case "down-to":
		err = goose.DownTo(sql, migrationDir, parseVersion())

	case "redo":
		err = goose.Redo(sql, migrationDir)

	case "reset":
		err = goose.Reset(sql, migrationDir)

	case "create":
		if len(os.Args) < 3 {
			log.Fatal("missing migration name")
		}

		migrationType := "sql"
		if err := goose.Create(nil, migrationDir, os.Args[2], migrationType); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Migration created.")
		return

	case "status":
		err = goose.Status(sql, migrationDir)

	case "version":
		version, err := goose.GetDBVersion(sql)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(version)
		return

	default:
		usage()
	}

	if err != nil {
		log.Fatal(err)
	}
}
