package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"standard/pkg/constants"
	"standard/pkg/utils"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

const (
	dir = "cmd/migration/migrations"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	db, err := utils.SetupPostgresConnection()
	if err != nil {
		logrus.WithFields(logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryMigration}).Panic(err.Error())
	}
	defer db.Close()

	command := args[0]
	switch command {
	case "up", "down", "redo", "status":
		err = goose.RunContext(context.Background(), command, db.DB, dir, args...)
	default:
		err = goose.RunContext(context.Background(), "help", db.DB, dir, args...)
	}
	if err != nil {
		log.Fatal(err)
	}

}

func usage() {
	fmt.Println("Usage: myapp [OPTIONS] COMMAND")
	fmt.Println("Options:")
	fmt.Println("  -h, --help		Show this help message")
	fmt.Println("Commands:")
	fmt.Println("  up			Migrate the database to the most recent version available")
	fmt.Println("  down			Roll back the version by 1")
	fmt.Println("  redo			Roll back the most recently applied migration, then run it again")
	fmt.Println("  status		Print the status of all migrations")
}
