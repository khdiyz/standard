package main

import (
	"bufio"
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

	migrate "github.com/rubenv/sql-migrate"
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

	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}

	command := args[0]
	switch command {
	case "up":
		n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	case "down":
		n, err := migrate.ExecMax(db.DB, "postgres", migrations, migrate.Down, 1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rolled back %d migrations!\n", n)
	case "redo":
		n, err := migrate.ExecMax(db.DB, "postgres", migrations, migrate.Down, 1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rolled back %d migrations!\n", n)
		n, err = migrate.ExecMax(db.DB, "postgres", migrations, migrate.Up, 1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Re-applied %d migrations!\n", n)
	case "status":
		records, err := migrate.GetMigrationRecords(db.DB, "postgres")
		if err != nil {
			log.Fatal(err)
		}
		if len(records) == 0 {
			fmt.Println("No migrations applied!")
			return
		}
		for _, record := range records {
			fmt.Printf("Migration %s applied at %s\n", record.Id, record.AppliedAt)
		}
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
		flags.Usage()
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

	templateContent := `-- +migrate Up

-- +migrate Down

`
	tmpl, err := template.New("migration").Parse(templateContent)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, nil)
}
