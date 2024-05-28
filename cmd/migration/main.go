package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"standard/pkg/constants"
	"standard/pkg/utils"
	"strings"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

const (
	dir        = "cmd/migration/migrations"
	timeFormat = "20060102150405"
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
	case "create":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter name of migration: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		name = strings.TrimSpace(name) // Remove any surrounding whitespace including newline characters
		if name == "" {
			fmt.Println("Please provide a name for the new migration")
			return
		}
		err = createMigrationFile(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created new migration %s\n", name)
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

func createMigrationFile(name string) error {
	timestamp := time.Now().Format(timeFormat)
	filename := fmt.Sprintf("%s_%s.sql", timestamp, name)
	filepath := filepath.Join(dir, filename)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	templateContent := `-- +goose Up

-- +goose Down

`
	tmpl, err := template.New("migration").Parse(templateContent)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, nil)
}
